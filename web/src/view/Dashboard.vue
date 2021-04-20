<template>
  <el-card class="box-card">
    <el-row :gutter="30">
      <el-col :span="14">
        <el-calendar v-model="value"> </el-calendar>
      </el-col>
      <el-col :span="10">
        <el-card class="box-card" shadow="hover" style="margin-bottom: 20px;">
          <div slot="header" class="clearfix">
            <span>系统信息</span>
          </div>
          <div class="text item" style="margin-bottom: 10px;">本地工作空间：{{ localSpacePath }}</div>
          <div class="text item" style="margin-bottom: 10px;">远端工作空间：{{ remoteSpacePath }}</div>
          <div class="text item" style="margin-bottom: 10px;">当前配置地址：{{ currentConfigFilePath }}</div>
          <div class="text item" style="margin-bottom: 10px;">当前版本：{{ currentZoomVersion }}</div>
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