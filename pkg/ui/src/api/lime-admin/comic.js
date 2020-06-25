import request from './request'

export function getComic(id) {
  return request({
    url: '/admin/comics/comic/' + id,
    method: 'GET'
  })
}

export function ComicList(query) {
  return request({
    url: '/admin/comics/comic',
    method: 'GET',
    params: query
  })
}

export function updateComic(id, data) {
  return request({
    url: '/admin/comics/comic/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/comics/comic/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createComic(data) {
  return request({
    url: '/admin/comics/comic',
    method: 'post',
    data: data
  })
}

export function deleteComic(id,data) {
  return request({
    url: '/admin/comics/comic/' + id,
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
