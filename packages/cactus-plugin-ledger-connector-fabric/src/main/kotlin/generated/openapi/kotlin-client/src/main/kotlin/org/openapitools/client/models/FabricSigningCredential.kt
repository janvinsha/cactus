/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models

import org.openapitools.client.models.FabricSigningCredentialType
import org.openapitools.client.models.VaultTransitKey
import org.openapitools.client.models.WebSocketKey

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * 
 *
 * @param keychainId 
 * @param keychainRef 
 * @param type 
 * @param vaultTransitKey 
 * @param webSocketKey 
 */


data class FabricSigningCredential (

    @Json(name = "keychainId")
    val keychainId: kotlin.String,

    @Json(name = "keychainRef")
    val keychainRef: kotlin.String,

    @Json(name = "type")
    val type: FabricSigningCredentialType? = null,

    @Json(name = "vaultTransitKey")
    val vaultTransitKey: VaultTransitKey? = null,

    @Json(name = "webSocketKey")
    val webSocketKey: WebSocketKey? = null

)

