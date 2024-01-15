package emobility

import (
	"testing"
	"time"

	"github.com/enbility/eebus-go/util"
	"github.com/enbility/spine-go/model"
	"github.com/stretchr/testify/assert"
)

func Test_EVGetTimeSlotConstraints(t *testing.T) {
	emobilty, eebusService := setupEmobility()

	constraints, err := emobilty.EVTimeSlotConstraints()
	assert.Equal(t, uint(0), constraints.MinSlots)
	assert.Equal(t, uint(0), constraints.MaxSlots)
	assert.Equal(t, time.Duration(0), constraints.MinSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.MaxSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.SlotDurationStepSize)
	assert.NotEqual(t, err, nil)

	localDevice, localEntity, remoteDevice, entites, _ := setupDevices(eebusService)
	emobilty.evseEntity = entites[0]
	emobilty.evEntity = entites[1]

	constraints, err = emobilty.EVTimeSlotConstraints()
	assert.Equal(t, uint(0), constraints.MinSlots)
	assert.Equal(t, uint(0), constraints.MaxSlots)
	assert.Equal(t, time.Duration(0), constraints.MinSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.MaxSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.SlotDurationStepSize)
	assert.NotEqual(t, err, nil)

	emobilty.evTimeSeries = timeSeriesConfiguration(localEntity, emobilty.evEntity)

	constraints, err = emobilty.EVTimeSlotConstraints()
	assert.Equal(t, uint(0), constraints.MinSlots)
	assert.Equal(t, uint(0), constraints.MaxSlots)
	assert.Equal(t, time.Duration(0), constraints.MinSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.MaxSlotDuration)
	assert.Equal(t, time.Duration(0), constraints.SlotDurationStepSize)
	assert.NotEqual(t, err, nil)

	datagram := datagramForEntityAndFeatures(false, localDevice, localEntity, emobilty.evEntity, model.FeatureTypeTypeTimeSeries, model.RoleTypeServer, model.RoleTypeClient)

	cmd := []model.CmdType{{
		TimeSeriesConstraintsListData: &model.TimeSeriesConstraintsListDataType{
			TimeSeriesConstraintsData: []model.TimeSeriesConstraintsDataType{
				{
					TimeSeriesId:         util.Ptr(model.TimeSeriesIdType(0)),
					SlotCountMin:         util.Ptr(model.TimeSeriesSlotCountType(1)),
					SlotCountMax:         util.Ptr(model.TimeSeriesSlotCountType(10)),
					SlotDurationMin:      model.NewDurationType(1 * time.Minute),
					SlotDurationMax:      model.NewDurationType(60 * time.Minute),
					SlotDurationStepSize: model.NewDurationType(1 * time.Minute),
				},
			},
		}}}

	datagram.Payload.Cmd = cmd

	err = localDevice.ProcessCmd(datagram, remoteDevice)
	assert.Nil(t, err)

	constraints, err = emobilty.EVTimeSlotConstraints()
	assert.Equal(t, uint(1), constraints.MinSlots)
	assert.Equal(t, uint(10), constraints.MaxSlots)
	assert.Equal(t, time.Duration(1*time.Minute), constraints.MinSlotDuration)
	assert.Equal(t, time.Duration(1*time.Hour), constraints.MaxSlotDuration)
	assert.Equal(t, time.Duration(1*time.Minute), constraints.SlotDurationStepSize)
	assert.Equal(t, err, nil)
}
