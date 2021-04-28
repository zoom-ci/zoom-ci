<template>
    <div class="layer-global">
        <div class="layer-header-box">
            <!-- <header class="layer-header" :style="{backgroundColor: bgColor}"> -->
            <header class="layer-header">
                <div class="header-left">
                    <router-link :to="{ path:'/dashboard' }" >
                        <svg class="svg-logo" version="1.1" xmlns="http://www.w3.org/2000/svg"
                            xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="130px" height="28px"
                            viewBox="0 0 160 36" enable-background="new 0 0 160 36" xml:space="preserve" style="margin-top:10px;">  <image
                            id="image0" width="160" height="36" x="0" y="0" href="data:image/png;base64,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tENOCgFBNKEVMUKB6zvqqHL0/T7zkk0WEcQJtCCaullCwMGNKtTIYekICFDvGuzt5qp0tMWbop8GT7F3aKSVFy+PJgHEjT/+nWSrRaUeW2O0oQBsHVVQ7ABWYyVyio9cXCQCU9TmlTp2G5QSEVHZ+r4SbumE6WD0oPkHlShS3sNB+BEyExMNG198to4GaQ9lEgZhQKmbogUVXn545Un1BT3LDUqJiBhEJBBHJTxwqrxk9HZ0qKo2HnGEEUQQkaTLiXGb2ostd+cKqls2m7oYZSVBacU1+Hv72KlZbG3izQWT1HEemQ6clEIgUmoUFaiAGw7ZoIjqPDw5miJ4JVgeIohbV1I1bF54isCmlMg81zf3O6ZyZMr0lggWPJalkpxuSZrawj8SuLxiqa3gMQhyQ/TE4t/cl26ALHny676JVCTKwB6CB264YIddNlQkqlpSIT7jiCCKqHTIn53uoWJzM9kokCObZB7GZbiuUwgIG540S8HR7IjqCEoYCUNmaUEfyxAUERFJiVcd2xL7oR90n+Viere1D8TQV29OfHVXLG1RlUAAzWhDjH2zRHLPdMAJpxy8wpYM3quICQ8eccSRQFwyGcpVQbXIlkBC5rPIRqTwOs4Mex4qm0AwY78QKTSfExyQ1nvya2Z48u+EU45Fqy4J6TydimpB4eD21L1w7YVjup/rPc1m1XulQKB8tKnlqy8+mr7xs3TbOAoREbTgFELU5JSZENmb1yY7fyuW8sokICR/NH8DxOKwEbYvATUWNbk9uZYkjHUWpWkjCoEdXtShEzywVXKdVFEJrNuL5IYYbVKHiUoXkzj8OAW/0XiScdspL59kWO/U/wu50WXCpvdh8H9DaCSAEDhRhTpUw50xDJUC0kyf/FJHLiplKvJIezyu9mVrT2hMDCglCYTAAaBwyWmlEhZl70HyI/Lnei7Svg/sP008UkizRaEiiSMAKm9g2dslJIM+pDk+hhACCGgs/8soATKMlqhAYnJIfi/ccCadnNpLXJS5nZftzUI6YS/LaHforlIJJ4fRdMMli0r7QpTt06OISVbqZTEpNRjqDCFydDF7SaYfZYcxr52OMgoJA72XqJW+9lzU0uR/5ZGkw+D/Af87EreCBemMAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDIxLTA0LTE1VDE5OjUwOjIxKzA4OjAwRQh6fwAAACV0RVh0ZGF0ZTptb2RpZnkAMjAyMS0wNC0xNVQxOTo1MDoyMSswODowMDRVwsMAAAAgdEVYdHNvZnR3YXJlAGh0dHBzOi8vaW1hZ2VtYWdpY2sub3JnvM8dnQAAABh0RVh0VGh1bWI6OkRvY3VtZW50OjpQYWdlcwAxp/+7LwAAABd0RVh0VGh1bWI6OkltYWdlOjpIZWlnaHQAODlc0vi3AAAAF3RFWHRUaHVtYjo6SW1hZ2U6OldpZHRoADI3NkgEdsYAAAAZdEVYdFRodW1iOjpNaW1ldHlwZQBpbWFnZS9wbmc/slZOAAAAF3RFWHRUaHVtYjo6TVRpbWUAMTYxODQ4NzQyMYlE+jwAAAASdEVYdFRodW1iOjpTaXplADg4MDdCQsMaxcEAAABGdEVYdFRodW1iOjpVUkkAZmlsZTovLy9hcHAvdG1wL2ltYWdlbGMvaW1ndmlldzJfOV8xNjA5OTAzNTMyMDE1NzUyMl83OF9bMF05kaPyAAAAAElFTkSuQmCC"></image>
                        </svg>
                    </router-link>
                </div>
                <div class="header-right">
                    <span class="r-item">
                        <el-dropdown trigger="click" @command="userSettingHandler">
                            <span class="item app-cursor">
                                {{ $t('say_hello_to_login_user', {username : $store.getters['account/getUserName']}) }}
                                <i class="iconfont small icon-arrow-down"></i>
                            </span>
                            <el-dropdown-menu slot="dropdown" class="app-header-dropdown">
                                <el-dropdown-item>{{$store.getters['account/getTrueName']}}</el-dropdown-item>
                                <el-dropdown-item command="setting"><i class="iconfont small left icon-user"></i>{{ $t('personal_setting') }}</el-dropdown-item>
                                <el-dropdown-item command="password"><i class="iconfont small left icon-key"></i>{{ $t('password') }}</el-dropdown-item>
                                <el-dropdown-item command="language"><i class="el-icon-sort"></i>{{ $t('language') }}</el-dropdown-item>
                                <a class="app-dropdown-link" href="https://zoom-ci.github.io/docs/" target="_blank">
                                    <el-dropdown-item divided><i class="iconfont small left icon-pull-request"></i>{{ $t('zoom_document') }}</el-dropdown-item>
                                </a>
                                <el-dropdown-item command="logout"><i class="iconfont small left icon-logout"></i>{{ $t('sign_out') }}</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </span>
                </div>
            </header>
        </div>
        <section class="layer-container">
            <aside class="layer-aside">
                <ScrollBar>
                    <el-menu class="aside-menu" :default-active="activeMenu" :router="true" :unique-opened="true">
                        <template v-for="menu in AppMenu">
                            <el-submenu
                                    v-if="menu.children && (menu.children.length > 1 || (menu.children.length == 1 && !menu.children[0].meta.single))"
                                    :index="menu.name" :key="menu.name">
                                <template slot="title">
                                    <span v-if="menu.meta.icon" class="iconfont left"
                                          :class="menu.meta.icon"></span><span>{{ menu.meta.title }}</span>
                                </template>
                                <template v-for="childMenu in menu.children">
                                    <el-menu-item v-if="!(childMenu.meta && childMenu.meta.hide)"
                                                  :route="{name: childMenu.name}" :index="childMenu.name"
                                                  :key="childMenu.name">
                                        <i class="iconfont small left">
                                            <svg viewBox="0 0 1024 1024" width="1em" height="1em" fill="currentColor"
                                                 aria-hidden="true">
                                                <path d="M384.023552 384.083968l256.016384 0 0 256.016384-256.016384 0 0-256.016384Z"></path>
                                            </svg>
                                        </i>
                                        <span>{{ childMenu.meta.title }}</span>
                                    </el-menu-item>
                                </template>
                            </el-submenu>
                            <el-menu-item :route="{name: menu.children[0].name}"
                                          v-else-if="menu.children && menu.children.length == 1"
                                          :index="menu.children[0].name" :key="menu.children[0].name">
                                <i v-if="menu.children[0].meta.icon" class="iconfont left"
                                   :class="menu.children[0].meta.icon"></i>
                                <span>{{ menu.children[0].meta.title }}</span>
                            </el-menu-item>
                        </template>
                    </el-menu>
                </ScrollBar>
            </aside>
            <main class="layer-main">
                <el-breadcrumb separator="/" class="bread-crumb">
                    <el-breadcrumb-item><i class="iconfont small icon-breadcrumbs"></i></el-breadcrumb-item>
                    <el-breadcrumb-item v-for="b in breadcrumb" :key="b">{{ b }}</el-breadcrumb-item>
                </el-breadcrumb>
                <div class="container">
                    <router-view/>
                    <div class="app-cpy">
                        <p>
                            ©️ {{ new Date().getFullYear() }} <a href="https://github.com/zoom-ci/zoom-ci/"
                                                                 target="_blank">Zoom-CI</a>. All Rights Reserved. MIT
                            License.
                        </p>
                    </div>
                </div>
            </main>
        </section>

        <el-dialog :width="$root.DialogSmallWidth" :title="$t('personal_setting')" :visible.sync="settingDialogVisible"
                   @close="closeUserSettingDialogHandler">
            <el-form class="app-form" ref="settingDialogRef" :model="settingForm" size="medium" label-width="80px">
                <el-form-item
                        :label="$t('role')">
                    {{ $store.getters['account/getRoleName'] }}
                </el-form-item>
                <el-form-item
                        :label="$t('username')">
                    <el-input :value="$store.getters['account/getUserName']" :disabled="true"></el-input>
                    <span class="app-input-help"><i class="el-icon-info"></i> {{ $t('contact_admin_for_change_username')}}</span>
                </el-form-item>
                <el-form-item
                        :label="$t('email')">
                    <el-input :value="$store.getters['account/getEmail']" :disabled="true"></el-input>
                    <span class="app-input-help"><i class="el-icon-info"></i> {{ $t('contact_admin_for_change_useremail')}}</span>
                </el-form-item>
                <el-form-item
                        :label="$t('truename')"
                        prop="truename">
                    <el-input v-model="settingForm.truename"></el-input>
                </el-form-item>
                <el-form-item
                        :label="$t('mobile')"
                        prop="mobile">
                    <el-input v-model="settingForm.mobile"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="small" @click="closeUserSettingDialogHandler">{{ $t('cancel')}}</el-button>
                <el-button :loading="btnLoading" size="small" type="primary" @click="dialogUserSettingSubmitHandler">{{
                    $t('enter')}}
                </el-button>
            </div>
        </el-dialog>

        <el-dialog :width="$root.DialogSmallWidth" :title="$t('change_password')" :visible.sync="passwordDialogVisible"
                   @close="closePasswordSettingDialogHandler">
            <el-form class="app-form" ref="passwordDialogRef" :model="passwordForm" size="medium" label-width="150px">
                <el-form-item
                        :label="$t('current_password')"
                        prop="password"
                        :rules="[
                    { min: 6, max: 20, message: this.$t('strlen_between', {min: 6, max: 20}), trigger: 'blur'},
                    { required: true, message: this.$t('current_password_cannot_empty'), trigger: 'blur'},
                ]">
                    <el-input type="password" :placeholder="$t('please_input_password_length_limit')"
                              v-model="passwordForm.password" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item
                        :label="$t('new_password')"
                        prop="new_password"
                        :rules="[
                    { min: 6, max: 20, message: this.$t('strlen_between', {min: 6, max: 20}), trigger: 'blur'},
                    { required: true, message: this.$t('new_password_cannot_empty'), trigger: 'blur'},
                ]">
                    <el-input type="password" :placeholder="$t('please_input_password_length_limit')"
                              v-model="passwordForm.new_password" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button size="small" @click="closePasswordSettingDialogHandler">{{ $t('cancel')}}</el-button>
                <el-button :loading="btnLoading" size="small" type="primary"
                           @click="dialogPasswordSettingSubmitHandler">{{ $t('enter')}}
                </el-button>
            </div>
        </el-dialog>

    </div>
</template>

<script>
    import ScrollBar from '@/component/ScrollBar';
    import {routerMap} from '@/router'
    import Code from '@/lib/code'
    import util from '@/lib/util'
    import {loginStatusApi, logoutApi} from '@/api/login'
    import {installStatusApi} from '@/api/system'
    import {userPasswordApi, userSettingApi} from '@/api/user'

    export default {
        data() {
            return {
                breadcrumb: [],
                activeMenu: '',
                btnLoading: false,
                settingDialogVisible: false,
                settingForm: {},
                passwordDialogVisible: false,
                passwordForm: {},
            }
        },
        computed: {
            AppMenu() {
                let menu = []
                routerMap.forEach(first => {
                    let newSecond = []
                    let newFirst = Object.assign({}, first)
                    first.children.forEach(second => {
                        if (!second.meta.role || this.$root.CheckPrivs(second.meta.role)) {
                            newSecond.push(second)
                        }
                    })
                    if (newSecond.length > 0) {
                        newFirst.children = newSecond
                        menu.push(newFirst)
                    }
                })
                return menu
            }
        },
        watch: {
            '$route.name'() {
                this.breadcrumbItems()
                this.initActiveMenu()
            },
            AppMenu() {
                this.breadcrumbItems()
            },
        },
        components: {
            ScrollBar,
        },
        methods: {
            userSettingHandler(cmd) {
                switch (cmd) {
                    case 'logout':
                        logoutApi().then(res => {
                            this.$router.push({name: 'login'})
                        })
                        break
                    case 'setting':
                        this.showUserSettingDialogHandler()
                        break
                    case 'password':
                        this.showPasswordSettingDialogHandler()
                        break
                    case 'language':
                        util.switchLanguage()
                        break
                }
            },

            showPasswordSettingDialogHandler() {
                this.passwordDialogVisible = true
            },
            closePasswordSettingDialogHandler() {
                this.passwordDialogVisible = false
            },
            dialogPasswordSettingSubmitHandler() {
                let vm = this
                this.$refs.passwordDialogRef.validate((valid) => {
                    if (!valid) {
                        return false;
                    }
                    let postData = {
                        password: this.$root.Md5Sum(vm.passwordForm.password),
                        new_password: this.$root.Md5Sum(vm.passwordForm.new_password)
                    }
                    this.btnLoading = true
                    userPasswordApi(postData).then(res => {
                        this.$root.MessageSuccess(() => {
                            this.btnLoading = false
                            this.closePasswordSettingDialogHandler()
                        })
                    }).catch(err => {
                        if (err.code == Code.CODE_ERR_USER_OR_PASS_WRONG) {
                            this.$message({
                                message: this.$t('current_password_error'),
                                type: 'warning',
                            });
                        }
                        this.btnLoading = false
                    })
                })
            },

            showUserSettingDialogHandler() {
                this.settingForm = {
                    truename: this.$store.getters['account/getTrueName'],
                    mobile: this.$store.getters['account/getMobile'],
                }
                this.settingDialogVisible = true
            },
            closeUserSettingDialogHandler() {
                this.settingDialogVisible = false
            },
            dialogUserSettingSubmitHandler() {
                this.btnLoading = true
                userSettingApi(this.settingForm).then(res => {
                    this.$root.MessageSuccess(() => {
                        this.btnLoading = false
                        this.closeUserSettingDialogHandler()
                    })
                    this.$store.dispatch('account/userSetting', {
                        mobile: this.settingForm.mobile,
                        truename: this.settingForm.truename,
                    })
                }).catch(err => {
                    this.btnLoading = false
                })
            },

            initActiveMenu() {
                this.activeMenu = this.$route.name
            },
            breadcrumbItems() {
                let breadcrumb = []
                this.AppMenu.forEach(menu => {
                    menu.children.forEach(sub => {
                        if (sub.name != this.$route.name) {
                            return
                        }
                        if (menu.meta.title) {
                            breadcrumb.push(menu.meta.title)
                        }
                        breadcrumb.push(sub.meta.title)
                    })
                })
                this.breadcrumb = breadcrumb
            },
            initInstallStatus() {
                let self = this;
                installStatusApi().then(res => {
                    if (res.is_installed) {
                        self.initLoginStatus()
                    } else {
                        this.$router.push({name: 'install'})
                    }
                }).catch(err => {
                    this.$router.push({name: 'install'})
                })
            },
            initLoginStatus() {
                loginStatusApi().then(res => {
                    if (res.is_login) {
                        this.$store.dispatch('account/status', {
                            user_id: res.user_id,
                            username: res.username,
                            email: res.email,
                            mobile: res.mobile,
                            privilege: res.privilege ? res.privilege : [],
                            role_name: res.role_name,
                            truename: res.truename,
                        })
                    } else {
                        this.$message.error(this.$t('not_login'), 1)
                        this.$router.push({name: 'login'})
                    }
                }).catch(err => {
                    this.$message.error(this.$t('not_login'), 1)
                    this.$router.push({name: 'login'})
                })
            },
        },
        mounted() {
            this.initInstallStatus()
            this.initActiveMenu()
        },
    }
</script>
