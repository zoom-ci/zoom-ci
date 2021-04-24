import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

import zh from './zh_cn.js'
import en from './en.js'
import util from '../lib/util.js'
const messages = {
    'zh-cn': zh,
    'en': en,
}

let localeLang
if (global.navigator.language) {
    localeLang = global.navigator.language
    localeLang = localeLang.toLowerCase()
}
if (localeLang.indexOf('en') != 0) {
    localeLang = 'zh-cn'
}
if (localeLang.indexOf('en') != 0 && util.GetLanguage() != 'en') {
    localeLang = 'zh-cn'
} else {
    localeLang = 'en'
}
util.SetLanguage(localeLang)
// localeLang = 'en'
const i18n = new VueI18n({
    locale: localeLang,
    messages,
})

export default i18n