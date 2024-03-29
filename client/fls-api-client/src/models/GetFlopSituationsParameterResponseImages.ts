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
/**
 * 
 * @export
 * @interface GetFlopSituationsParameterResponseImages
 */
export interface GetFlopSituationsParameterResponseImages {
    /**
     * 
     * @type {string}
     * @memberof GetFlopSituationsParameterResponseImages
     */
    url: string;
    /**
     * 
     * @type {string}
     * @memberof GetFlopSituationsParameterResponseImages
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof GetFlopSituationsParameterResponseImages
     */
    description: string;
}

export function GetFlopSituationsParameterResponseImagesFromJSON(json: any): GetFlopSituationsParameterResponseImages {
    return GetFlopSituationsParameterResponseImagesFromJSONTyped(json, false);
}

export function GetFlopSituationsParameterResponseImagesFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetFlopSituationsParameterResponseImages {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'url': json['url'],
        'name': json['name'],
        'description': json['description'],
    };
}

export function GetFlopSituationsParameterResponseImagesToJSON(value?: GetFlopSituationsParameterResponseImages | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'url': value.url,
        'name': value.name,
        'description': value.description,
    };
}


