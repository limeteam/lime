import request from './request'

export function categoryList() {
  return request({
    url: '/comics/categories',
    method: 'GET'
  })
}

export function updateCategory(id, data) {
  return request({
    url: '/comics/categories/' + id,
    method: 'put',
    data: data
  })
}

export function createCategory(data) {
  return request({
    url: '/comics/categories',
    method: 'post',
    data: data
  })
}

export function deleteCategory(id,data) {
  return request({
    url: '/comics/categories/' + id,
    method: 'delete',
    data: data
  })
}
