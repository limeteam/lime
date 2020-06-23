import request from './request'

export function getChapter(id) {
  return request({
    url: '/admin/comics/chapters/' + id,
    method: 'GET'
  })
}

export function ChapterList(query) {
  return request({
    url: '/admin/comics/chapters',
    method: 'GET',
    params: query
  })
}

export function updateChapter(id, data) {
  return request({
    url: '/admin/comics/chapters/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/admin/comics//chapters/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createChapter(data) {
  return request({
    url: '/admin/comics/chapters',
    method: 'post',
    data: data
  })
}

export function deleteChapter(id,data) {
  return request({
    url: '/admin/comics/chapters/' + id,
    method: 'delete',
    data: data
  })
}
