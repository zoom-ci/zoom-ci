<template>
  <div class="app-install">
    <div
      class="install-logo"
      style="text-align: center; justify-content: center; margin-top: 10vh"
    >
      <img style="height: 60px" src="../asset/logo.png" />
    </div>
    <div class="app-install-inner">
      <div class="install-container">
        <el-card class="install-box">
          <div class="install-title">快速安装</div>
          <el-form
            @keyup.enter.native="installHandler"
            ref="installFormRef"
            :model="installForm"
            :rules="installRules"
            size="medium"
            class="install-form"
            label-position="right"
            label-width="150px"
          >
            <el-form-item prop="mysql_host" label="Mysql Host" required>
              <el-input
                v-model="installForm.mysql_host"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item prop="mysql_port" label="Mysql Port" required>
              <el-input
                v-model="installForm.mysql_port"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item prop="mysql_username" label="Mysql Username" required>
              <el-input
                v-model="installForm.mysql_username"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item prop="mysql_password" label="Mysql Password" required>
              <el-input
                v-model="installForm.mysql_password"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item prop="mysql_dbname" label="Mysql Dbname" required>
              <el-input
                v-model="installForm.mysql_dbname"
                autocomplete="off"
              ></el-input>
            </el-form-item>
            <el-form-item>
              <el-button @click="installHandler" type="primary">{{
                $t("install")
              }}</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
      <div class="install-cpy">
        ©️ {{ new Date().getFullYear() }}
        <a href="https://github.com/zoom-ci/zoom-ci/" target="_blank">Zoom</a>.
        All Rights Reserved. MIT License.
      </div>
    </div>
  </div>
</template>

<script>
import { installApi, installStatusApi } from "@/api/system";
import Code from "@/lib/code";
export default {
  data() {
    return {
      installLoadding: false,
      installForm: {
        mysql_host: "127.0.0.1",
        mysql_port: "3306",
        mysql_username: "zoom",
        mysql_password: "zoom",
        mysql_dbname: "zoom",
      },
      installRules: {
        mysql_host: [
          {
            required: true,
            message: this.$t("please_input_mysql_host"),
            trigger: "blur",
          },
        ],
        mysql_port: [
          {
            required: true,
            message: this.$t("please_input_mysql_port"),
            trigger: "blur",
          },
        ],
        mysql_username: [
          {
            required: true,
            message: this.$t("please_input_mysql_username"),
            trigger: "blur",
          },
        ],
        mysql_password: [
          {
            required: true,
            message: this.$t("please_input_mysql_password"),
            trigger: "blur",
          },
        ],
        mysql_dbname: [
          {
            required: true,
            message: this.$t("please_input_mysql_dbname"),
            trigger: "blur",
          },
        ],
      },
    };
  },
  computed: {},
  mounted() {
    this.initInstallStatus();
  },
  methods: {
    initInstallStatus() {
      let self = this;
      installStatusApi()
        .then((res) => {
          if (res.is_installed) {
            this.$router.push({ name: "login" });
          }
        })
        .catch((err) => {
          //
        });
    },
    installHandler() {
      this.$refs.installFormRef.validate((valid) => {
        if (!valid) {
          return false;
        }
        let postData = {
          mysql_host: this.installForm.mysql_host,
          mysql_port: this.installForm.mysql_port,
          mysql_username: this.installForm.mysql_username,
          mysql_password: this.installForm.mysql_password,
          mysql_dbname: this.installForm.mysql_dbname,
        };
        installApi(postData)
          .then((res) => {
            if (res.is_installed) {
              this.$message.success("安装成功");
              this.$router.push({ name: "login" });
            } else {
              this.$message.error("安装失败");
            }
          })
          .catch((err) => {
            if (err.code && err.code == Code.CODE_ERR_INSTALL_FAILED) {
              this.$message.error("安装失败, 错误信息: " + err.message);
            }
          });
      });
    },
  },
};
</script>

<style lang="scss" scope>
.app-install {
  width: 100%;
  height: 100%;
  position: relative;
  overflow: auto;
  min-height: 100%;
  background: #f0f2f5 url(../asset/background.svg) no-repeat 50%;
  background-size: 100%;
  .install-container {
    display: flex;
    justify-content: center;
    .install-box {
      margin-top: 5vh;
      width: 50vw;
      .install-title {
        font-weight: 500;
        text-align: center;
        font-size: 16px;
        margin-bottom: 20px;
      }
    }
  }
  .install-cpy {
    display: flex;
    justify-content: center;
    margin-top: 30px;
    color: #fff;
    a {
      margin: 0 5px;
      color: #fff;
    }
  }
}
</style>