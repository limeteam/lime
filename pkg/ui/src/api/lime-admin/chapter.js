import request from './request'
import qs from 'qs'

export function getChapter(id) {
  return request({
    url: 'chapters/' + id,
    method: 'GET'
  })
}

export function ChapterList() {
  return request({
    url: 'chapters',
    method: 'GET'
  })
}

export function updateChapter(id, data) {
  return request({
    url: 'chapters/' + id,
    method: 'put',
    data: data
  })
}

export function updatestatus(id, data) {
  return request({
    url: '/chapters/' + id + '/status',
    method: 'post',
    data: data
  })
}

export function createChapter(data) {
  return request({
    url: 'chapters',
    method: 'post',
    data: data
  })
}

export function deleteChapter(id,data) {
  return request({
    url: 'chapters/' + id,
    method: 'delete',
    data: data
  })
}