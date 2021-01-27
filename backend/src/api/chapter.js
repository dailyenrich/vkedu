import axios from "../lib/ajax"
import backend_uri from "./backend-api"

class Course {
    constructor() {

    }

    getOneByID(id) {
        return new Promise(function(resolve, reject) {
            axios.get(backend_uri.CHAPTER_ONE_BY_ID_URI.replace(":id", id)).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    edit(cource) {
        return new Promise(function(resolve, reject) {
            axios.put(backend_uri.CHAPTER_UPDATE_URI.replace(":id", cource.id), cource).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    delete(id) {
        return new Promise(function(resolve, reject) {
            axios.delete(backend_uri.CHAPTER_DELETE_URI.replace(":id", id)).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    store(chapter) {
        return new Promise(function(resolve, reject) {
            axios.post(backend_uri.CHAPTER_STORE_URI, chapter).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    list(condition) {
        return new Promise(function(resolve, reject) {
            axios.get(backend_uri.CHAPTER_LIST_URI, condition).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }
}

export default new Course