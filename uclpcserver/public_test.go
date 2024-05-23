package uclpcserver

import (
	"time"

	"github.com/enbility/cemd/api"
	"github.com/enbility/cemd/util"
	eebusutil "github.com/enbility/eebus-go/util"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
)

func (s *UCLPCServerSuite) Test_ConsumptionLimit() {
	limit, err := s.sut.ConsumptionLimit()
	assert.Equal(s.T(), 0.0, limit.Value)
	assert.NotNil(s.T(), err)

	newLimit := api.LoadLimit{
		Duration:     time.Duration(time.Hour * 2),
		IsActive:     true,
		IsChangeable: true,
		Value:        16,
	}
	err = s.sut.SetConsumptionLimit(newLimit)
	assert.Nil(s.T(), err)

	limit, err = s.sut.ConsumptionLimit()
	assert.Equal(s.T(), 16.0, limit.Value)
	assert.Nil(s.T(), err)
}

func (s *UCLPCServerSuite) Test_PendingConsumptionLimits() {
	data := s.sut.PendingConsumptionLimits()
	assert.Equal(s.T(), 0, len(data))

	msgCounter := model.MsgCounterType(500)

	msg := &spineapi.Message{
		RequestHeader: &model.HeaderType{
			MsgCounter: eebusutil.Ptr(msgCounter),
		},
		Cmd: model.CmdType{
			LoadControlLimitListData: &model.LoadControlLimitListDataType{
				LoadControlLimitData: []model.LoadControlLimitDataType{
					{
						LimitId:       eebusutil.Ptr(model.LoadControlLimitIdType(0)),
						IsLimitActive: eebusutil.Ptr(true),
						Value:         model.NewScaledNumberType(1000),
						TimePeriod:    model.NewTimePeriodTypeWithRelativeEndTime(time.Minute * 2),
					},
				},
			},
		},
		DeviceRemote: s.remoteDevice,
		EntityRemote: s.monitoredEntity,
	}

	s.sut.loadControlWriteCB(msg)

	data = s.sut.PendingConsumptionLimits()
	assert.Equal(s.T(), 1, len(data))

	s.sut.ApproveOrDenyConsumptionLimit(model.MsgCounterType(499), true, "")

	s.sut.ApproveOrDenyConsumptionLimit(msgCounter, false, "leave me alone")
}

func (s *UCLPCServerSuite) Test_Failsafe() {
	limit, changeable, err := s.sut.FailsafeConsumptionActivePowerLimit()
	assert.Equal(s.T(), 0.0, limit)
	assert.Equal(s.T(), false, changeable)
	assert.NotNil(s.T(), err)

	err = s.sut.SetFailsafeConsumptionActivePowerLimit(10, true)
	assert.Nil(s.T(), err)

	limit, changeable, err = s.sut.FailsafeConsumptionActivePowerLimit()
	assert.Equal(s.T(), 10.0, limit)
	assert.Equal(s.T(), true, changeable)
	assert.Nil(s.T(), err)

	// The actual tests of the functionality is located in the util package
	duration, changeable, err := s.sut.FailsafeDurationMinimum()
	assert.Equal(s.T(), time.Duration(0), duration)
	assert.Equal(s.T(), false, changeable)
	assert.NotNil(s.T(), err)

	err = s.sut.SetFailsafeDurationMinimum(time.Duration(time.Hour*1), true)
	assert.NotNil(s.T(), err)

	err = s.sut.SetFailsafeDurationMinimum(time.Duration(time.Hour*2), true)
	assert.Nil(s.T(), err)

	limit, changeable, err = s.sut.FailsafeConsumptionActivePowerLimit()
	assert.Equal(s.T(), 10.0, limit)
	assert.Equal(s.T(), true, changeable)
	assert.Nil(s.T(), err)

	duration, changeable, err = s.sut.FailsafeDurationMinimum()
	assert.Equal(s.T(), time.Duration(time.Hour*2), duration)
	assert.Equal(s.T(), true, changeable)
	assert.Nil(s.T(), err)
}

func (s *UCLPCServerSuite) Test_IsHeartbeatWithinDuration() {
	assert.Nil(s.T(), s.sut.heartbeatDiag)

	value := s.sut.IsHeartbeatWithinDuration()
	assert.False(s.T(), value)

	remoteDiagServer := s.monitoredEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	assert.NotNil(s.T(), remoteDiagServer)

	var err error
	s.sut.heartbeatDiag, err = util.DeviceDiagnosis(s.service, s.monitoredEntity)
	assert.NotNil(s.T(), remoteDiagServer)
	assert.Nil(s.T(), err)

	// add heartbeat data to the remoteDiagServer
	timestamp := time.Now().Add(-time.Second * 121)
	data := &model.DeviceDiagnosisHeartbeatDataType{
		Timestamp:        model.NewAbsoluteOrRelativeTimeTypeFromTime(timestamp),
		HeartbeatCounter: eebusutil.Ptr(uint64(1)),
		HeartbeatTimeout: model.NewDurationType(time.Second * 120),
	}
	err1 := remoteDiagServer.UpdateData(model.FunctionTypeDeviceDiagnosisHeartbeatData, data, nil, nil)
	assert.Nil(s.T(), err1)

	value = s.sut.IsHeartbeatWithinDuration()
	assert.False(s.T(), value)

	timestamp = time.Now()
	data.Timestamp = model.NewAbsoluteOrRelativeTimeTypeFromTime(timestamp)

	err1 = remoteDiagServer.UpdateData(model.FunctionTypeDeviceDiagnosisHeartbeatData, data, nil, nil)
	assert.Nil(s.T(), err1)

	value = s.sut.IsHeartbeatWithinDuration()
	assert.True(s.T(), value)
}

func (s *UCLPCServerSuite) Test_ContractualConsumptionNominalMax() {
	value, err := s.sut.ContractualConsumptionNominalMax()
	assert.Equal(s.T(), 0.0, value)
	assert.NotNil(s.T(), err)

	err = s.sut.SetContractualConsumptionNominalMax(10)
	assert.Nil(s.T(), err)

	value, err = s.sut.ContractualConsumptionNominalMax()
	assert.Equal(s.T(), 10.0, value)
	assert.Nil(s.T(), err)
}
