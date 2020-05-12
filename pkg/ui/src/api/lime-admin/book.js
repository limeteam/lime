import request from './request'
import qs from 'qs'

export function getBook(id) {
  return request({
    url: 'books/' + id,
    method: 'GET'
  })
}

export function BookList() {
  return request({
    url: 'books',
    method: 'GET'
  })
}

export function updateBook(id, data) {
  return request({
    url: 'books/' + id,
    method: 'put',
    data: data
  })
}

export function createBook(data) {
  return request({
    url: 'books',
    method: 'post',
    data: data
  })
}

export function deleteBook(id,data) {
  return request({
    url: 'books/' + id,
    method: 'delete',
    data: data
  })
}