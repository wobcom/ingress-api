// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var MessageFieldPathsNested = []string{
	"Payload",
	"Payload.join_accept_payload",
	"Payload.join_accept_payload.cf_list",
	"Payload.join_accept_payload.cf_list.ch_masks",
	"Payload.join_accept_payload.cf_list.freq",
	"Payload.join_accept_payload.cf_list.type",
	"Payload.join_accept_payload.dev_addr",
	"Payload.join_accept_payload.dl_settings",
	"Payload.join_accept_payload.dl_settings.opt_neg",
	"Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"Payload.join_accept_payload.dl_settings.rx2_dr",
	"Payload.join_accept_payload.encrypted",
	"Payload.join_accept_payload.join_nonce",
	"Payload.join_accept_payload.net_id",
	"Payload.join_accept_payload.rx_delay",
	"Payload.join_request_payload",
	"Payload.join_request_payload.dev_eui",
	"Payload.join_request_payload.dev_nonce",
	"Payload.join_request_payload.join_eui",
	"Payload.mac_payload",
	"Payload.mac_payload.decoded_payload",
	"Payload.mac_payload.f_hdr",
	"Payload.mac_payload.f_hdr.dev_addr",
	"Payload.mac_payload.f_hdr.f_cnt",
	"Payload.mac_payload.f_hdr.f_ctrl",
	"Payload.mac_payload.f_hdr.f_ctrl.ack",
	"Payload.mac_payload.f_hdr.f_ctrl.adr",
	"Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"Payload.mac_payload.f_hdr.f_opts",
	"Payload.mac_payload.f_port",
	"Payload.mac_payload.frm_payload",
	"Payload.mac_payload.full_f_cnt",
	"Payload.rejoin_request_payload",
	"Payload.rejoin_request_payload.dev_eui",
	"Payload.rejoin_request_payload.join_eui",
	"Payload.rejoin_request_payload.net_id",
	"Payload.rejoin_request_payload.rejoin_cnt",
	"Payload.rejoin_request_payload.rejoin_type",
	"m_hdr",
	"m_hdr.m_type",
	"m_hdr.major",
	"mic",
}

var MessageFieldPathsTopLevel = []string{
	"Payload",
	"m_hdr",
	"mic",
}
var MHDRFieldPathsNested = []string{
	"m_type",
	"major",
}

var MHDRFieldPathsTopLevel = []string{
	"m_type",
	"major",
}
var MACPayloadFieldPathsNested = []string{
	"decoded_payload",
	"f_hdr",
	"f_hdr.dev_addr",
	"f_hdr.f_cnt",
	"f_hdr.f_ctrl",
	"f_hdr.f_ctrl.ack",
	"f_hdr.f_ctrl.adr",
	"f_hdr.f_ctrl.adr_ack_req",
	"f_hdr.f_ctrl.class_b",
	"f_hdr.f_ctrl.f_pending",
	"f_hdr.f_opts",
	"f_port",
	"frm_payload",
	"full_f_cnt",
}

var MACPayloadFieldPathsTopLevel = []string{
	"decoded_payload",
	"f_hdr",
	"f_port",
	"frm_payload",
	"full_f_cnt",
}
var FHDRFieldPathsNested = []string{
	"dev_addr",
	"f_cnt",
	"f_ctrl",
	"f_ctrl.ack",
	"f_ctrl.adr",
	"f_ctrl.adr_ack_req",
	"f_ctrl.class_b",
	"f_ctrl.f_pending",
	"f_opts",
}

var FHDRFieldPathsTopLevel = []string{
	"dev_addr",
	"f_cnt",
	"f_ctrl",
	"f_opts",
}
var FCtrlFieldPathsNested = []string{
	"ack",
	"adr",
	"adr_ack_req",
	"class_b",
	"f_pending",
}

var FCtrlFieldPathsTopLevel = []string{
	"ack",
	"adr",
	"adr_ack_req",
	"class_b",
	"f_pending",
}
var JoinRequestPayloadFieldPathsNested = []string{
	"dev_eui",
	"dev_nonce",
	"join_eui",
}

var JoinRequestPayloadFieldPathsTopLevel = []string{
	"dev_eui",
	"dev_nonce",
	"join_eui",
}
var RejoinRequestPayloadFieldPathsNested = []string{
	"dev_eui",
	"join_eui",
	"net_id",
	"rejoin_cnt",
	"rejoin_type",
}

var RejoinRequestPayloadFieldPathsTopLevel = []string{
	"dev_eui",
	"join_eui",
	"net_id",
	"rejoin_cnt",
	"rejoin_type",
}
var JoinAcceptPayloadFieldPathsNested = []string{
	"cf_list",
	"cf_list.ch_masks",
	"cf_list.freq",
	"cf_list.type",
	"dev_addr",
	"dl_settings",
	"dl_settings.opt_neg",
	"dl_settings.rx1_dr_offset",
	"dl_settings.rx2_dr",
	"encrypted",
	"join_nonce",
	"net_id",
	"rx_delay",
}

var JoinAcceptPayloadFieldPathsTopLevel = []string{
	"cf_list",
	"dev_addr",
	"dl_settings",
	"encrypted",
	"join_nonce",
	"net_id",
	"rx_delay",
}
var DLSettingsFieldPathsNested = []string{
	"opt_neg",
	"rx1_dr_offset",
	"rx2_dr",
}

var DLSettingsFieldPathsTopLevel = []string{
	"opt_neg",
	"rx1_dr_offset",
	"rx2_dr",
}
var CFListFieldPathsNested = []string{
	"ch_masks",
	"freq",
	"type",
}

var CFListFieldPathsTopLevel = []string{
	"ch_masks",
	"freq",
	"type",
}
var LoRaDataRateFieldPathsNested = []string{
	"bandwidth",
	"spreading_factor",
}

var LoRaDataRateFieldPathsTopLevel = []string{
	"bandwidth",
	"spreading_factor",
}
var FSKDataRateFieldPathsNested = []string{
	"bit_rate",
}

var FSKDataRateFieldPathsTopLevel = []string{
	"bit_rate",
}
var DataRateFieldPathsNested = []string{
	"modulation",
	"modulation.fsk",
	"modulation.fsk.bit_rate",
	"modulation.lora",
	"modulation.lora.bandwidth",
	"modulation.lora.spreading_factor",
}

var DataRateFieldPathsTopLevel = []string{
	"modulation",
}
var TxSettingsFieldPathsNested = []string{
	"coding_rate",
	"data_rate",
	"data_rate.modulation",
	"data_rate.modulation.fsk",
	"data_rate.modulation.fsk.bit_rate",
	"data_rate.modulation.lora",
	"data_rate.modulation.lora.bandwidth",
	"data_rate.modulation.lora.spreading_factor",
	"data_rate_index",
	"downlink",
	"downlink.antenna_index",
	"downlink.invert_polarization",
	"downlink.tx_power",
	"enable_crc",
	"frequency",
	"time",
	"timestamp",
}

var TxSettingsFieldPathsTopLevel = []string{
	"coding_rate",
	"data_rate",
	"data_rate_index",
	"downlink",
	"enable_crc",
	"frequency",
	"time",
	"timestamp",
}
var GatewayAntennaIdentifiersFieldPathsNested = []string{
	"antenna_index",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var GatewayAntennaIdentifiersFieldPathsTopLevel = []string{
	"antenna_index",
	"gateway_ids",
}
var UplinkTokenFieldPathsNested = []string{
	"concentrator_time",
	"ids",
	"ids.antenna_index",
	"ids.gateway_ids",
	"ids.gateway_ids.eui",
	"ids.gateway_ids.gateway_id",
	"server_time",
	"timestamp",
}

var UplinkTokenFieldPathsTopLevel = []string{
	"concentrator_time",
	"ids",
	"server_time",
	"timestamp",
}
var DownlinkPathFieldPathsNested = []string{
	"path",
	"path.fixed",
	"path.fixed.antenna_index",
	"path.fixed.gateway_ids",
	"path.fixed.gateway_ids.eui",
	"path.fixed.gateway_ids.gateway_id",
	"path.uplink_token",
}

var DownlinkPathFieldPathsTopLevel = []string{
	"path",
}
var TxRequestFieldPathsNested = []string{
	"absolute_time",
	"advanced",
	"class",
	"downlink_paths",
	"frequency_plan_id",
	"priority",
	"rx1_data_rate_index",
	"rx1_delay",
	"rx1_frequency",
	"rx2_data_rate_index",
	"rx2_frequency",
}

var TxRequestFieldPathsTopLevel = []string{
	"absolute_time",
	"advanced",
	"class",
	"downlink_paths",
	"frequency_plan_id",
	"priority",
	"rx1_data_rate_index",
	"rx1_delay",
	"rx1_frequency",
	"rx2_data_rate_index",
	"rx2_frequency",
}
var MACCommandFieldPathsNested = []string{
	"cid",
	"payload",
	"payload.adr_param_setup_req",
	"payload.adr_param_setup_req.adr_ack_delay_exponent",
	"payload.adr_param_setup_req.adr_ack_limit_exponent",
	"payload.beacon_freq_ans",
	"payload.beacon_freq_ans.frequency_ack",
	"payload.beacon_freq_req",
	"payload.beacon_freq_req.frequency",
	"payload.beacon_timing_ans",
	"payload.beacon_timing_ans.channel_index",
	"payload.beacon_timing_ans.delay",
	"payload.dev_status_ans",
	"payload.dev_status_ans.battery",
	"payload.dev_status_ans.margin",
	"payload.device_mode_conf",
	"payload.device_mode_conf.class",
	"payload.device_mode_ind",
	"payload.device_mode_ind.class",
	"payload.device_time_ans",
	"payload.device_time_ans.time",
	"payload.dl_channel_ans",
	"payload.dl_channel_ans.channel_index_ack",
	"payload.dl_channel_ans.frequency_ack",
	"payload.dl_channel_req",
	"payload.dl_channel_req.channel_index",
	"payload.dl_channel_req.frequency",
	"payload.duty_cycle_req",
	"payload.duty_cycle_req.max_duty_cycle",
	"payload.force_rejoin_req",
	"payload.force_rejoin_req.data_rate_index",
	"payload.force_rejoin_req.max_retries",
	"payload.force_rejoin_req.period_exponent",
	"payload.force_rejoin_req.rejoin_type",
	"payload.link_adr_ans",
	"payload.link_adr_ans.channel_mask_ack",
	"payload.link_adr_ans.data_rate_index_ack",
	"payload.link_adr_ans.tx_power_index_ack",
	"payload.link_adr_req",
	"payload.link_adr_req.channel_mask",
	"payload.link_adr_req.channel_mask_control",
	"payload.link_adr_req.data_rate_index",
	"payload.link_adr_req.nb_trans",
	"payload.link_adr_req.tx_power_index",
	"payload.link_check_ans",
	"payload.link_check_ans.gateway_count",
	"payload.link_check_ans.margin",
	"payload.new_channel_ans",
	"payload.new_channel_ans.data_rate_ack",
	"payload.new_channel_ans.frequency_ack",
	"payload.new_channel_req",
	"payload.new_channel_req.channel_index",
	"payload.new_channel_req.frequency",
	"payload.new_channel_req.max_data_rate_index",
	"payload.new_channel_req.min_data_rate_index",
	"payload.ping_slot_channel_ans",
	"payload.ping_slot_channel_ans.data_rate_index_ack",
	"payload.ping_slot_channel_ans.frequency_ack",
	"payload.ping_slot_channel_req",
	"payload.ping_slot_channel_req.data_rate_index",
	"payload.ping_slot_channel_req.frequency",
	"payload.ping_slot_info_req",
	"payload.ping_slot_info_req.period",
	"payload.raw_payload",
	"payload.rejoin_param_setup_ans",
	"payload.rejoin_param_setup_ans.max_time_exponent_ack",
	"payload.rejoin_param_setup_req",
	"payload.rejoin_param_setup_req.max_count_exponent",
	"payload.rejoin_param_setup_req.max_time_exponent",
	"payload.rekey_conf",
	"payload.rekey_conf.minor_version",
	"payload.rekey_ind",
	"payload.rekey_ind.minor_version",
	"payload.reset_conf",
	"payload.reset_conf.minor_version",
	"payload.reset_ind",
	"payload.reset_ind.minor_version",
	"payload.rx_param_setup_ans",
	"payload.rx_param_setup_ans.rx1_data_rate_offset_ack",
	"payload.rx_param_setup_ans.rx2_data_rate_index_ack",
	"payload.rx_param_setup_ans.rx2_frequency_ack",
	"payload.rx_param_setup_req",
	"payload.rx_param_setup_req.rx1_data_rate_offset",
	"payload.rx_param_setup_req.rx2_data_rate_index",
	"payload.rx_param_setup_req.rx2_frequency",
	"payload.rx_timing_setup_req",
	"payload.rx_timing_setup_req.delay",
	"payload.tx_param_setup_req",
	"payload.tx_param_setup_req.downlink_dwell_time",
	"payload.tx_param_setup_req.max_eirp_index",
	"payload.tx_param_setup_req.uplink_dwell_time",
}

var MACCommandFieldPathsTopLevel = []string{
	"cid",
	"payload",
}
var FrequencyValueFieldPathsNested = []string{
	"value",
}

var FrequencyValueFieldPathsTopLevel = []string{
	"value",
}
var DataRateOffsetValueFieldPathsNested = []string{
	"value",
}

var DataRateOffsetValueFieldPathsTopLevel = []string{
	"value",
}
var DataRateIndexValueFieldPathsNested = []string{
	"value",
}

var DataRateIndexValueFieldPathsTopLevel = []string{
	"value",
}
var PingSlotPeriodValueFieldPathsNested = []string{
	"value",
}

var PingSlotPeriodValueFieldPathsTopLevel = []string{
	"value",
}
var AggregatedDutyCycleValueFieldPathsNested = []string{
	"value",
}

var AggregatedDutyCycleValueFieldPathsTopLevel = []string{
	"value",
}
var RxDelayValueFieldPathsNested = []string{
	"value",
}

var RxDelayValueFieldPathsTopLevel = []string{
	"value",
}
var ADRAckLimitExponentValueFieldPathsNested = []string{
	"value",
}

var ADRAckLimitExponentValueFieldPathsTopLevel = []string{
	"value",
}
var ADRAckDelayExponentValueFieldPathsNested = []string{
	"value",
}

var ADRAckDelayExponentValueFieldPathsTopLevel = []string{
	"value",
}
var TxSettings_DownlinkFieldPathsNested = []string{
	"antenna_index",
	"invert_polarization",
	"tx_power",
}

var TxSettings_DownlinkFieldPathsTopLevel = []string{
	"antenna_index",
	"invert_polarization",
	"tx_power",
}
var MACCommand_ResetIndFieldPathsNested = []string{
	"minor_version",
}

var MACCommand_ResetIndFieldPathsTopLevel = []string{
	"minor_version",
}
var MACCommand_ResetConfFieldPathsNested = []string{
	"minor_version",
}

var MACCommand_ResetConfFieldPathsTopLevel = []string{
	"minor_version",
}
var MACCommand_LinkCheckAnsFieldPathsNested = []string{
	"gateway_count",
	"margin",
}

var MACCommand_LinkCheckAnsFieldPathsTopLevel = []string{
	"gateway_count",
	"margin",
}
var MACCommand_LinkADRReqFieldPathsNested = []string{
	"channel_mask",
	"channel_mask_control",
	"data_rate_index",
	"nb_trans",
	"tx_power_index",
}

var MACCommand_LinkADRReqFieldPathsTopLevel = []string{
	"channel_mask",
	"channel_mask_control",
	"data_rate_index",
	"nb_trans",
	"tx_power_index",
}
var MACCommand_LinkADRAnsFieldPathsNested = []string{
	"channel_mask_ack",
	"data_rate_index_ack",
	"tx_power_index_ack",
}

var MACCommand_LinkADRAnsFieldPathsTopLevel = []string{
	"channel_mask_ack",
	"data_rate_index_ack",
	"tx_power_index_ack",
}
var MACCommand_DutyCycleReqFieldPathsNested = []string{
	"max_duty_cycle",
}

var MACCommand_DutyCycleReqFieldPathsTopLevel = []string{
	"max_duty_cycle",
}
var MACCommand_RxParamSetupReqFieldPathsNested = []string{
	"rx1_data_rate_offset",
	"rx2_data_rate_index",
	"rx2_frequency",
}

var MACCommand_RxParamSetupReqFieldPathsTopLevel = []string{
	"rx1_data_rate_offset",
	"rx2_data_rate_index",
	"rx2_frequency",
}
var MACCommand_RxParamSetupAnsFieldPathsNested = []string{
	"rx1_data_rate_offset_ack",
	"rx2_data_rate_index_ack",
	"rx2_frequency_ack",
}

var MACCommand_RxParamSetupAnsFieldPathsTopLevel = []string{
	"rx1_data_rate_offset_ack",
	"rx2_data_rate_index_ack",
	"rx2_frequency_ack",
}
var MACCommand_DevStatusAnsFieldPathsNested = []string{
	"battery",
	"margin",
}

var MACCommand_DevStatusAnsFieldPathsTopLevel = []string{
	"battery",
	"margin",
}
var MACCommand_NewChannelReqFieldPathsNested = []string{
	"channel_index",
	"frequency",
	"max_data_rate_index",
	"min_data_rate_index",
}

var MACCommand_NewChannelReqFieldPathsTopLevel = []string{
	"channel_index",
	"frequency",
	"max_data_rate_index",
	"min_data_rate_index",
}
var MACCommand_NewChannelAnsFieldPathsNested = []string{
	"data_rate_ack",
	"frequency_ack",
}

var MACCommand_NewChannelAnsFieldPathsTopLevel = []string{
	"data_rate_ack",
	"frequency_ack",
}
var MACCommand_DLChannelReqFieldPathsNested = []string{
	"channel_index",
	"frequency",
}

var MACCommand_DLChannelReqFieldPathsTopLevel = []string{
	"channel_index",
	"frequency",
}
var MACCommand_DLChannelAnsFieldPathsNested = []string{
	"channel_index_ack",
	"frequency_ack",
}

var MACCommand_DLChannelAnsFieldPathsTopLevel = []string{
	"channel_index_ack",
	"frequency_ack",
}
var MACCommand_RxTimingSetupReqFieldPathsNested = []string{
	"delay",
}

var MACCommand_RxTimingSetupReqFieldPathsTopLevel = []string{
	"delay",
}
var MACCommand_TxParamSetupReqFieldPathsNested = []string{
	"downlink_dwell_time",
	"max_eirp_index",
	"uplink_dwell_time",
}

var MACCommand_TxParamSetupReqFieldPathsTopLevel = []string{
	"downlink_dwell_time",
	"max_eirp_index",
	"uplink_dwell_time",
}
var MACCommand_RekeyIndFieldPathsNested = []string{
	"minor_version",
}

var MACCommand_RekeyIndFieldPathsTopLevel = []string{
	"minor_version",
}
var MACCommand_RekeyConfFieldPathsNested = []string{
	"minor_version",
}

var MACCommand_RekeyConfFieldPathsTopLevel = []string{
	"minor_version",
}
var MACCommand_ADRParamSetupReqFieldPathsNested = []string{
	"adr_ack_delay_exponent",
	"adr_ack_limit_exponent",
}

var MACCommand_ADRParamSetupReqFieldPathsTopLevel = []string{
	"adr_ack_delay_exponent",
	"adr_ack_limit_exponent",
}
var MACCommand_DeviceTimeAnsFieldPathsNested = []string{
	"time",
}

var MACCommand_DeviceTimeAnsFieldPathsTopLevel = []string{
	"time",
}
var MACCommand_ForceRejoinReqFieldPathsNested = []string{
	"data_rate_index",
	"max_retries",
	"period_exponent",
	"rejoin_type",
}

var MACCommand_ForceRejoinReqFieldPathsTopLevel = []string{
	"data_rate_index",
	"max_retries",
	"period_exponent",
	"rejoin_type",
}
var MACCommand_RejoinParamSetupReqFieldPathsNested = []string{
	"max_count_exponent",
	"max_time_exponent",
}

var MACCommand_RejoinParamSetupReqFieldPathsTopLevel = []string{
	"max_count_exponent",
	"max_time_exponent",
}
var MACCommand_RejoinParamSetupAnsFieldPathsNested = []string{
	"max_time_exponent_ack",
}

var MACCommand_RejoinParamSetupAnsFieldPathsTopLevel = []string{
	"max_time_exponent_ack",
}
var MACCommand_PingSlotInfoReqFieldPathsNested = []string{
	"period",
}

var MACCommand_PingSlotInfoReqFieldPathsTopLevel = []string{
	"period",
}
var MACCommand_PingSlotChannelReqFieldPathsNested = []string{
	"data_rate_index",
	"frequency",
}

var MACCommand_PingSlotChannelReqFieldPathsTopLevel = []string{
	"data_rate_index",
	"frequency",
}
var MACCommand_PingSlotChannelAnsFieldPathsNested = []string{
	"data_rate_index_ack",
	"frequency_ack",
}

var MACCommand_PingSlotChannelAnsFieldPathsTopLevel = []string{
	"data_rate_index_ack",
	"frequency_ack",
}
var MACCommand_BeaconTimingAnsFieldPathsNested = []string{
	"channel_index",
	"delay",
}

var MACCommand_BeaconTimingAnsFieldPathsTopLevel = []string{
	"channel_index",
	"delay",
}
var MACCommand_BeaconFreqReqFieldPathsNested = []string{
	"frequency",
}

var MACCommand_BeaconFreqReqFieldPathsTopLevel = []string{
	"frequency",
}
var MACCommand_BeaconFreqAnsFieldPathsNested = []string{
	"frequency_ack",
}

var MACCommand_BeaconFreqAnsFieldPathsTopLevel = []string{
	"frequency_ack",
}
var MACCommand_DeviceModeIndFieldPathsNested = []string{
	"class",
}

var MACCommand_DeviceModeIndFieldPathsTopLevel = []string{
	"class",
}
var MACCommand_DeviceModeConfFieldPathsNested = []string{
	"class",
}

var MACCommand_DeviceModeConfFieldPathsTopLevel = []string{
	"class",
}
