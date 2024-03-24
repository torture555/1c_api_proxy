export const HOST_NAME = "localhost"
export const PORT_BACKEND = "10001"
export const URL_BACKEND = "http://" + HOST_NAME + ":" + PORT_BACKEND

export const EnumStatusConnection = [
    { key: true, value: "Есть соединение", color: "#008000" },
    { key: false, value: "Неизвестно", color: "#D3D3D3" }
]

export const MAP_METHODS = [
    { key: "SetDBParams" , value: "/db/set" },
    { key: "GetDBParams" , value: "/db/get" },
    { key: "GetDBStatus", value: "/db/status" },
    { key: "GetInfobases", value: "/infobase" },
    { key: "AddInfobase", value: "/infobase/add" },
    { key: "EditInfobase", value: "/infobase/edit" },
    { key: "DeleteInfobase", value: "/infobase/delete" },
    { key: "StatusInfobase", value: "/infobase/status" },
    { key: "ReloadConnection", value: "/infobase/reload" },
    { key: "GetLogs", value: "/log/get" },
] // how to use: const reformattedArray = MAP_METHODS.map(({ key, value }) => ({ [key]: value }));

export const EnumLogsLevel = [
    { key: "Info", value: "Info", color: "#87CEEB" },
    { key: "Warn", value: "Warn", color: "#FFD700" },
    { key: "Error", value: "Error",color: "#8B0000" },
]