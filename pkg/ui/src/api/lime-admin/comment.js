import request from './request'

export function CommentList() {
  return request({
    url: '/admin/novels/comments',
    method: 'GET'
  })
}

export function deleteComment(id,data) {
  return request({
    url: '/admin/novels/comments/' + id,
    method: 'delete',
    data: data
  })
}
