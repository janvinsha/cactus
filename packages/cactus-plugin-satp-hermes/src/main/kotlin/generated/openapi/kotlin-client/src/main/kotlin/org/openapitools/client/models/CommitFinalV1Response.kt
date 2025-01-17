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


import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * 
 *
 * @param sessionID 
 * @param messageType 
 * @param clientIdentityPubkey 
 * @param serverIdentityPubkey 
 * @param commitAcknowledgementClaim 
 * @param hashCommitFinal 
 * @param signature 
 * @param sequenceNumber 
 * @param commitAcknowledgementClaimFormat 
 * @param serverTransferNumber 
 */


data class CommitFinalV1Response (

    @Json(name = "sessionID")
    val sessionID: kotlin.String,

    @Json(name = "messageType")
    val messageType: kotlin.String,

    @Json(name = "clientIdentityPubkey")
    val clientIdentityPubkey: kotlin.String,

    @Json(name = "serverIdentityPubkey")
    val serverIdentityPubkey: kotlin.String,

    @Json(name = "commitAcknowledgementClaim")
    val commitAcknowledgementClaim: kotlin.String,

    @Json(name = "hashCommitFinal")
    val hashCommitFinal: kotlin.String,

    @Json(name = "signature")
    val signature: kotlin.String,

    @Json(name = "sequenceNumber")
    val sequenceNumber: java.math.BigDecimal,

    @Json(name = "commitAcknowledgementClaimFormat")
    val commitAcknowledgementClaimFormat: kotlin.Any? = null,

    @Json(name = "serverTransferNumber")
    val serverTransferNumber: kotlin.Int? = null

)

