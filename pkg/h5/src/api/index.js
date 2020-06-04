import request from '@/utils/request'


// 获取某本书
export function getBook(id) {
    return request({
        url: '/v1/books/' + id,
        method: 'GET'
    })
}

// 获取章节列表
export function getTitleList(book_id) {
    return request({
        url: '/v1/chapters',
        method: 'GET',
        params: {"book_id": book_id}
    })
}

// 获取书本
export function getBookContent(book_id, chapter_id) {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        // params: query
    })
}



// 获取书本列表
export function getBookList() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}


// 获取某一类
export function getBookType() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}


// 获取书架
export function getShelf() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}


// 加入书架

export function postShelf() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 检查是否在书架
export function checkShelf() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 搜索
export function search() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 获取最近阅读记录
export function getLate() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 添加到最近阅读记录
export function postLately() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 登录
export function login() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}

// 注册
export function register() {
    return request({
        url: '/v1/chapters/' + book_id,
        method: 'GET',
        params: query
    })
}
