import request from './request'
import qs from 'qs'

export function categoryList() {
  return request({
    url: 'categories',
    method: 'GET'
  })
}

export function updateCategory(id, data) {
  return request({
    url: 'categories/' + id,
    method: 'put',
    data: data
  })
}

export function createCategory(data) {
  return request({
    url: 'categories',
    method: 'post',
    data: data
  })
}

export function deleteCategory(id,data) {
  return request({
    url: 'categories/' + id,
    method: 'delete',
    data: data
  })
}