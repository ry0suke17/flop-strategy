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

/**
 * 
 * @export
 * @enum {string}
 */
export enum PotType {
    PotTypeSRP = 'PotTypeSRP',
    PotType3Bet = 'PotType_3_Bet',
    PotType4Bet = 'PotType_4_Bet'
}

export function PotTypeFromJSON(json: any): PotType {
    return PotTypeFromJSONTyped(json, false);
}

export function PotTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PotType {
    return json as PotType;
}

export function PotTypeToJSON(value?: PotType | null): any {
    return value as any;
}

