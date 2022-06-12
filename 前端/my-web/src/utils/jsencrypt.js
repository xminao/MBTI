import JSEncrypt from 'jsencrypt/bin/jsencrypt'

const publicKey = 'MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALKKhM8s0+DdBQGCuaGCUFn0Z3k5fgRJjOKo1qzmPuZUioPN+oao1Ec8mdT72HCesf+Lr8DpVE0ql9wYKj782fsCAwEAAQ=='
const privateKey = 'MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAsoqEzyzT4N0FAYK5oYJQWfRneTl+BEmM4qjWrOY+5lSKg836hqjURzyZ1PvYcJ6x/4uvwOlUTSqX3BgqPvzZ+wIDAQABAkAL6cxE2rPknDraR/PFACIGbpV89EYhWRsKgMOaU968lajwFXGRww7784tWKS1a4Qe/TnEw1zsmKd9Kpoqeipf5AiEA5mGX/OLmbYzx2Et0k0RK+zWscvtkMZzCUp90tf+CoRUCIQDGZSjW35wwV62OcrnRyRPPHza6JmF8DuQs2mqYTeyyzwIgRH0A7RobPLgo8Y9FAb7Mt6+2K5yWpbUfg+aJkM177R0CIAgYRIxP312RlkFDqRVIMQi4a1E5E60uJl02JDdGpqu9AiEAzBoSgpQjGrfoMmGwQ4dW//vExqvevpLyfJKe6XwLRHw='  //存放生成的私钥

// 加密
export function encrypt(txt) {
  const encryptor = new JSEncrypt()
  encryptor.setPublicKey(publicKey) // 设置公钥
  return encryptor.encrypt(txt) // 对需要加密的数据进行加密
}

// 解密
export function decrypt(txt) {
  const encryptor = new JSEncrypt()
  encryptor.setPrivateKey(privateKey)
  return encryptor.decrypt(txt)
}
