/* tslint:disable */
/* eslint-disable */
/**
 * flop-strategy
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    GetFlopSituationsParameterResponseImages,
    GetFlopSituationsParameterResponseImagesFromJSON,
    GetFlopSituationsParameterResponseImagesFromJSONTyped,
    GetFlopSituationsParameterResponseImagesToJSON,
    PlayerPositionType,
    PlayerPositionTypeFromJSON,
    PlayerPositionTypeFromJSONTyped,
    PlayerPositionTypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface GetFlopSituationsParameterResponse
 */
export interface GetFlopSituationsParameterResponse {
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    ipBetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    oopBetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    ipCheckFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    oopCheckFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    ip33BetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    oop33BetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    ip67BetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    oop67BetFreq: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    ipEquity: number;
    /**
     * 
     * @type {number}
     * @memberof GetFlopSituationsParameterResponse
     */
    oopEquity: number;
    /**
     * 
     * @type {PlayerPositionType}
     * @memberof GetFlopSituationsParameterResponse
     */
    heroPositionType: PlayerPositionType;
    /**
     * 
     * @type {Array<GetFlopSituationsParameterResponseImages>}
     * @memberof GetFlopSituationsParameterResponse
     */
    images?: Array<GetFlopSituationsParameterResponseImages>;
}

export function GetFlopSituationsParameterResponseFromJSON(json: any): GetFlopSituationsParameterResponse {
    return GetFlopSituationsParameterResponseFromJSONTyped(json, false);
}

export function GetFlopSituationsParameterResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetFlopSituationsParameterResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ipBetFreq': json['ip_bet_freq'],
        'oopBetFreq': json['oop_bet_freq'],
        'ipCheckFreq': json['ip_check_freq'],
        'oopCheckFreq': json['oop_check_freq'],
        'ip33BetFreq': json['ip_33_bet_freq'],
        'oop33BetFreq': json['oop_33_bet_freq'],
        'ip67BetFreq': json['ip_67_bet_freq'],
        'oop67BetFreq': json['oop_67_bet_freq'],
        'ipEquity': json['ip_equity'],
        'oopEquity': json['oop_equity'],
        'heroPositionType': PlayerPositionTypeFromJSON(json['hero_position_type']),
        'images': !exists(json, 'images') ? undefined : ((json['images'] as Array<any>).map(GetFlopSituationsParameterResponseImagesFromJSON)),
    };
}

export function GetFlopSituationsParameterResponseToJSON(value?: GetFlopSituationsParameterResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ip_bet_freq': value.ipBetFreq,
        'oop_bet_freq': value.oopBetFreq,
        'ip_check_freq': value.ipCheckFreq,
        'oop_check_freq': value.oopCheckFreq,
        'ip_33_bet_freq': value.ip33BetFreq,
        'oop_33_bet_freq': value.oop33BetFreq,
        'ip_67_bet_freq': value.ip67BetFreq,
        'oop_67_bet_freq': value.oop67BetFreq,
        'ip_equity': value.ipEquity,
        'oop_equity': value.oopEquity,
        'hero_position_type': PlayerPositionTypeToJSON(value.heroPositionType),
        'images': value.images === undefined ? undefined : ((value.images as Array<any>).map(GetFlopSituationsParameterResponseImagesToJSON)),
    };
}


