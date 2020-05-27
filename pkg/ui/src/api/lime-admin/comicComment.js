import request from './request'

export function CommentList() {
  return request({
    url: '/admin/comics/comments',
    method: 'GET'
  })
}

export function deleteComment(id,data) {
  return request({
    url: '/admin/comics/comments/' + id,
    method: 'delete',
    data: data
  })
}
