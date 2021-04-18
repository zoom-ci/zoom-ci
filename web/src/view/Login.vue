<template>
    <div class="app-login">
        <div class="login-logo" style="text-align:center;justify-content: center;margin-top:20vh;"><img style="height:60px;" src="../asset/logo.png"></div>
        <div class="app-login-inner">
            <div class="login-container">
                <el-card class="login-box">
                    <div class="login-title">{{ $t('welcome_to_login_zoom') }}</div>
                    <el-form @keyup.enter.native="loginHandler" ref="loginFormRef" :model="loginForm" :rules="loginRules" size="medium" class="login-form">
                        <el-form-item prop="username">
                            <el-input :placeholder="$t('username_or_email')" prefix-icon="iconfont icon-user" v-model="loginForm.username" autocomplete="off"></el-input>
                        </el-form-item>
                        <el-form-item prop="password">
                            <el-input :placeholder="$t('password')" prefix-icon="iconfont icon-lock" type="password" v-model="loginForm.password" autocomplete="off"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button @click="loginHandler" type="primary" style="width: 100%;">{{ $t('login') }}</el-button>
                        </el-form-item>
                    </el-form>
                </el-card>
            </div>
            <div class="login-cpy">
                ©️ {{ new Date().getFullYear() }} <a href="https://github.com/zoom-ci/zoom-ci/" target="_blank">Zoom</a>. All Rights Reserved. MIT License.
            </div>
        </div>
    </div>
</template>

<script>
import { loginApi } from '@/api/login'
import { installStatusApi } from '@/api/system'
import Code from '@/lib/code'
export default {
    data () {
        return {
            loginLoadding: false,
            loginForm: {
                username: "",
                password: "",
            },
            loginRules: {
                username: [{ required: true, message: this.$t('please_input_loginname'), trigger: 'blur' }],
                password: [
                    { required: true, message: this.$t('please_input_password'), trigger: 'blur' },
                    { min: 6, max: 20, message: this.$t('strlen_between', {min: 6, max: 20}), trigger: 'blur' },
                ],
            },
        }
    },
    computed: {
        
    },
    mounted() {
        this.initInstallStatus()
        console.log("xxxxx")
    },
    methods: {
        initInstallStatus() {
            let self = this;
            installStatusApi().then(res => {
                console.log(res)
                if (!res.is_installed) {
                    this.$router.push({name: 'install'})
                }
            }).catch(err => {
                this.$router.push({name: 'install'})
            })
        },
        loginHandler() {
            this.$refs.loginFormRef.validate((valid) => {
                if (!valid) {
                    return false
                }
                let postData = {
                    username: this.loginForm.username,
                    password: this.$root.Md5Sum(this.loginForm.password),
                }
                loginApi(postData).then(res => {
                    this.$root.SetLoginToken(res.token)
                    this.$router.push({name: 'dashboard'})
                }).catch(err => {
                    if (err.code && err.code == Code.CODE_ERR_LOGIN_FAILED) {
                        this.$message.error('登录失败, 错误信息: ' + err.message);
                    }
                })
            });
        },
    },
}
</script>

<style lang="scss" scope>
.app-login {
    width: 100%;
    height: 100%;
    position: relative;
    overflow: auto;
    min-height: 100%;
    background: #f0f2f5 url(../asset/background.svg) no-repeat 50%;
    background-size: 100%;
    .login-container {
        display: flex;
        justify-content: center;
        .login-box {
            margin-top: 5vh;
            width: 30vw;
            .login-title {
                font-weight: 500;
                text-align: center;
                font-size: 14px;
                margin-bottom: 20px;
            }
        }
    }
    .login-cpy {
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