package const_values

const FRONTEND_IP = "127.0.0.1"
const FRONTEND_PORT = ":8080"
const FRONTEND_DOMAIN_NAME = "http://" + FRONTEND_IP + FRONTEND_PORT
const FRONTEND_ALL_DOMAIN = "*"
const BACKEND_IP = "127.0.0.1"
const BACKEND_PORT = ":8080"
const BACKEND_SOURCE = "http://" + BACKEND_IP + BACKEND_PORT
const LOCALHOST = "http://localhost:8080"
const DATABASE = "dusha"
const PASSWORD = "xin365118"
const MYSQLCONNECTION = "root:" + PASSWORD + "@tcp(127.0.0.1:3306)/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
const BIGGEST_LIST_NUM = 8
