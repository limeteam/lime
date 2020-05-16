import request from './request'
import qs from 'qs'

export function getBook(id) {
  return request({
    url: 'books/' + id,
    method: 'GET'
  })
}

export function BookList(query) {
  return request({
    url: 'books',
    method: 'GET',
    params: query
  })
}

export function updateBook(id, data) {
  return request({
    url: 'books/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/books/' + id + '/status',
    method: 'post',
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

export function uploadcover(formData) {
  return request({
    url: 'upload/cover',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
     }
    // data: data
  })
}

export function uploadbookfile(formData) {
  return request({
    url: 'upload/book_file',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
     }
  })
}