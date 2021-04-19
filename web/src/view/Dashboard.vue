<template>
  <el-card class="box-card">
    <el-row :gutter="30">
      <el-col :span="16">
        <el-calendar v-model="value"> </el-calendar>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card" shadow="hover">
          <div slot="header" class="clearfix">
            <span>项目状态</span>
          </div>
          <div class="text item">上线中构建成功：N [查看]</div>
          <div class="text item">上线中构建失败：N [查看]</div>
          <div class="text item">上线成功的数量：N</div>
        </el-card>
        <el-card class="box-card" shadow="hover" style="margin-top:20px">
          <div slot="header" class="clearfix">
            <span>系统信息</span>
          </div>
          <div class="text item">本地工作空间：（已占用：XX GB）</div>
          <div class="text item">远端工作空间：</div>
          <div class="text item">当前配置地址：</div>
          <div class="text item">当前版本：</div>
        </el-card>
      </el-col>
    </el-row>
  </el-card>
</template>
<style>
.el-calendar-table .el-calendar-day {
    height:50px;
}
</style>
<script>
export default {
  data() {
    return {
      value: new Date(),
    };
  },
};
</script>