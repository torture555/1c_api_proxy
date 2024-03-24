import {EnumLogsLevel, EnumStatusConnection, MAP_METHODS} from "@/constants.js";

export function GetEnumValue(name, keyFind) {

    switch (name) {
        case "Methods":
            return MAP_METHODS.find((pair) => (pair.key === keyFind)).value
        case "StatusConnection":
            return EnumStatusConnection.find((pair) => (pair.key === keyFind)).value
        case "LogsLevel":
            return EnumLogsLevel.find((pair) => (pair.key === keyFind)).value
        default:
            return undefined
    }

}

export function GetEnumColor(name, keyFind) {

    switch (name) {
        case "StatusConnection":
            return EnumStatusConnection.find((pair) => (pair.key === keyFind)).color
        case "LogsLevel":
            return EnumLogsLevel.find((pair) => (pair.key === keyFind)).color
        default:
            return undefined
    }

}