import request from './request'

export function getBook(id) {
  return request({
    url: '/admin/novels/books/' + id,
    method: 'GET'
  })
}

export function BookList(query) {
  return request({
    url: '/admin/novels/books',
    method: 'GET',
    params: query
  })
}

export function updateBook(id, data) {
  return request({
    url: '/admin/novels/books/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/novels/books/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createBook(data) {
  return request({
    url: '/admin/novels/books',
    method: 'post',
    data: data
  })
}

export function deleteBook(id,data) {
  return request({
    url: '/admin/novels/books/' + id,
    method: 'delete',
    data: data
  })
}

export function uploadcover(formData) {
  return request({
    url: '/admin/upload/cover',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
     }
  })
}

export function uploadbookfile(formData) {
  return request({
    url: '/admin/upload/book_file',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
     }
  })
}
