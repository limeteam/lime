/**
 * 输出一个 UUID 字符串
 * Created by jiachenpan on 17/3/8.
 */
export default function createUniqueString() {
  const timestamp = +new Date() + ''
  const randomNum = parseInt((1 + Math.random()) * 65536) + ''
  return (+(randomNum + timestamp)).toString(32)
}
