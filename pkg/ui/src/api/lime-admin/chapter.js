import request from './request'
import qs from 'qs'

export function getChapter(id) {
  return request({
    url: '/novels/chapters/' + id,
    method: 'GET'
  })
}

export function ChapterList(query) {
  return request({
    url: '/novels/chapters',
    method: 'GET',
    params: query
  })
}

export function updateChapter(id, data) {
  return request({
    url: '/novels/chapters/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/novels//chapters/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createChapter(data) {
  return request({
    url: '/novels/chapters',
    method: 'post',
    data: data
  })
}

export function deleteChapter(id,data) {
  return request({
    url: '/novels/chapters/' + id,
    method: 'delete',
    data: data
  })
}
