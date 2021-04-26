<template>
  <el-card class="box-card">
    <el-row :gutter="30">
      <el-col :span="16">
        <el-calendar v-model="value"> </el-calendar>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card" shadow="hover" style="margin-bottom: 20px;">
          <div slot="header" class="clearfix">
            <span>{{ $t('system_info') }}</span>
          </div>
          <el-tooltip effect="dark" :content="$t('current_config_file_path') + ': ' + currentConfigFilePath" placement="top">          
            <div class="text item" style="margin-bottom: 10px;">{{ $t('local_space') }}：{{ localSpacePath }}</div>
          </el-tooltip>
          <div class="text item" style="margin-bottom: 10px;">{{ $t('remote_space') }}：{{ remoteSpacePath }}</div>
          <div class="text item" style="margin-bottom: 10px;">{{ $t('version') }}：{{ currentZoomVersion }}</div>
        </el-card>
      </el-col>
    </el-row>
  </el-card>
</template>
<style>
.el-calendar-table .el-calendar-day {
  height: 50px;
}
</style>
<script>
import { systemStatus } from "@/api/system";

export default {
  data() {
    return {
      value: new Date(),
      localSpacePath: "",
      remoteSpacePath: "",
      currentConfigFilePath: "",
      currentZoomVersion: "",
    };
  },
  mounted() {
    this.loadSystemStatus();
  },
  methods: {
    loadSystemStatus() {
      let self = this;
      systemStatus().then((res) => {
        if (res) {
          self.localSpacePath = res.localSpacePath;
          self.remoteSpacePath = res.remoteSpacePath;
          self.currentConfigFilePath = res.currentConfigFilePath;
          self.currentZoomVersion = res.currentZoomVersion;
        }
      });
    },
  },
};
</script>