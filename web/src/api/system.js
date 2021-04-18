import {get, post} from '@/lib/fetch.js'

export function installApi(data) {
    return post('system/install', data)
}
export function installStatusApi() {
    return get("/system/install_status")
}