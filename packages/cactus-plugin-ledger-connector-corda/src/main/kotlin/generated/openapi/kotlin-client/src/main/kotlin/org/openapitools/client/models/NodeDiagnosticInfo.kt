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

import org.openapitools.client.models.CordappInfo

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * A NodeDiagnosticInfo holds information about the current node version.
 *
 * @param cordapps A list of CorDapps currently installed on this node
 * @param platformVersion The platform version of this node. This number represents a released API version, and should be used to make functionality decisions (e.g. enabling an app feature only if an underlying platform feature exists)
 * @param revision The git commit hash this node was built from
 * @param vendor The vendor of this node
 * @param version The current node version string, e.g. 4.3, 4.4-SNAPSHOT. Note that this string is effectively freeform, and so should only be used for providing diagnostic information. It should not be used to make functionality decisions (the platformVersion is a better fit for this).
 */


data class NodeDiagnosticInfo (

    /* A list of CorDapps currently installed on this node */
    @Json(name = "cordapps")
    val cordapps: kotlin.collections.List<CordappInfo>,

    /* The platform version of this node. This number represents a released API version, and should be used to make functionality decisions (e.g. enabling an app feature only if an underlying platform feature exists) */
    @Json(name = "platformVersion")
    val platformVersion: kotlin.Int,

    /* The git commit hash this node was built from */
    @Json(name = "revision")
    val revision: kotlin.String,

    /* The vendor of this node */
    @Json(name = "vendor")
    val vendor: kotlin.String,

    /* The current node version string, e.g. 4.3, 4.4-SNAPSHOT. Note that this string is effectively freeform, and so should only be used for providing diagnostic information. It should not be used to make functionality decisions (the platformVersion is a better fit for this). */
    @Json(name = "version")
    val version: kotlin.String

)

