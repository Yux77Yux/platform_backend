
�x
google/api/http.proto
google.api"y
Http*
rules (2.google.api.HttpRuleRrulesE
fully_decode_reserved_expansion (RfullyDecodeReservedExpansion"�
HttpRule
selector (	Rselector
get (	H Rget
put (	H Rput
post (	H Rpost
delete (	H Rdelete
patch (	H Rpatch7
custom (2.google.api.CustomHttpPatternH Rcustom
body (	Rbody#
response_body (	RresponseBodyE
additional_bindings (2.google.api.HttpRuleRadditionalBindingsB	
pattern";
CustomHttpPattern
kind (	Rkind
path (	RpathB�
com.google.apiB	HttpProtoPZpgithub.com/Yux77Yux/platform_backend/generated/google.golang.org/genproto/googleapis/api/annotations;annotations��GAPIJ�s
 �
�
 2� Copyright 2024 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.


 

 
	
 
	
 �


 �

 "
	

 "

 *
	
 *

 '
	
 '

 "
	
$ "
�
  )� Defines the HTTP configuration for an API service. It contains a list of
 [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
 to one or more HTTP REST API methods.



 
�
   � A list of HTTP configuration rules that apply to individual API methods.

 **NOTE:** All service configuration rules follow "last one wins" order.


   


   

   

   
�
 (+� When set to true, URL path parameters will be fully URI-decoded except in
 cases of single segment matches in reserved expansion, where "%2F" will be
 left encoded.

 The default behavior is to not decode RFC 6570 reserved characters in multi
 segment matches.


 (

 (&

 ()*
�S
� ��R gRPC Transcoding

 gRPC Transcoding is a feature for mapping between a gRPC method and one or
 more HTTP REST endpoints. It allows developers to build a single API service
 that supports both gRPC APIs and REST APIs. Many systems, including [Google
 APIs](https://github.com/googleapis/googleapis),
 [Cloud Endpoints](https://cloud.google.com/endpoints), [gRPC
 Gateway](https://github.com/grpc-ecosystem/grpc-gateway),
 and [Envoy](https://github.com/envoyproxy/envoy) proxy support this feature
 and use it for large scale production services.

 `HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
 how different portions of the gRPC request message are mapped to the URL
 path, URL query parameters, and HTTP request body. It also controls how the
 gRPC response message is mapped to the HTTP response body. `HttpRule` is
 typically specified as an `google.api.http` annotation on the gRPC method.

 Each mapping specifies a URL path template and an HTTP method. The path
 template may refer to one or more fields in the gRPC request message, as long
 as each field is a non-repeated field with a primitive (non-message) type.
 The path template controls how fields of the request message are mapped to
 the URL path.

 Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get: "/v1/{name=messages/*}"
         };
       }
     }
     message GetMessageRequest {
       string name = 1; // Mapped to URL path.
     }
     message Message {
       string text = 1; // The resource content.
     }

 This enables an HTTP REST to gRPC mapping as below:

 - HTTP: `GET /v1/messages/123456`
 - gRPC: `GetMessage(name: "messages/123456")`

 Any fields in the request message which are not bound by the path template
 automatically become HTTP query parameters if there is no HTTP request body.
 For example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
             get:"/v1/messages/{message_id}"
         };
       }
     }
     message GetMessageRequest {
       message SubMessage {
         string subfield = 1;
       }
       string message_id = 1; // Mapped to URL path.
       int64 revision = 2;    // Mapped to URL query parameter `revision`.
       SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
     }

 This enables a HTTP JSON to RPC mapping as below:

 - HTTP: `GET /v1/messages/123456?revision=2&sub.subfield=foo`
 - gRPC: `GetMessage(message_id: "123456" revision: 2 sub:
 SubMessage(subfield: "foo"))`

 Note that fields which are mapped to URL query parameters must have a
 primitive type or a repeated primitive type or a non-repeated message type.
 In the case of a repeated type, the parameter can be repeated in the URL
 as `...?param=A&param=B`. In the case of a message type, each field of the
 message is mapped to a separate parameter, such as
 `...?foo.a=A&foo.b=B&foo.c=C`.

 For HTTP methods that allow a request body, the `body` field
 specifies the mapping. Consider a REST update method on the
 message resource collection:

     service Messaging {
       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "message"
         };
       }
     }
     message UpdateMessageRequest {
       string message_id = 1; // mapped to the URL
       Message message = 2;   // mapped to the body
     }

 The following HTTP JSON to RPC mapping is enabled, where the
 representation of the JSON in the request body is determined by
 protos JSON encoding:

 - HTTP: `PATCH /v1/messages/123456 { "text": "Hi!" }`
 - gRPC: `UpdateMessage(message_id: "123456" message { text: "Hi!" })`

 The special name `*` can be used in the body mapping to define that
 every field not bound by the path template should be mapped to the
 request body.  This enables the following alternative definition of
 the update method:

     service Messaging {
       rpc UpdateMessage(Message) returns (Message) {
         option (google.api.http) = {
           patch: "/v1/messages/{message_id}"
           body: "*"
         };
       }
     }
     message Message {
       string message_id = 1;
       string text = 2;
     }


 The following HTTP JSON to RPC mapping is enabled:

 - HTTP: `PATCH /v1/messages/123456 { "text": "Hi!" }`
 - gRPC: `UpdateMessage(message_id: "123456" text: "Hi!")`

 Note that when using `*` in the body mapping, it is not possible to
 have HTTP parameters, as all fields not bound by the path end in
 the body. This makes this option more rarely used in practice when
 defining REST APIs. The common usage of `*` is in custom methods
 which don't use the URL at all for transferring data.

 It is possible to define multiple HTTP methods for one RPC by using
 the `additional_bindings` option. Example:

     service Messaging {
       rpc GetMessage(GetMessageRequest) returns (Message) {
         option (google.api.http) = {
           get: "/v1/messages/{message_id}"
           additional_bindings {
             get: "/v1/users/{user_id}/messages/{message_id}"
           }
         };
       }
     }
     message GetMessageRequest {
       string message_id = 1;
       string user_id = 2;
     }

 This enables the following two alternative HTTP JSON to RPC mappings:

 - HTTP: `GET /v1/messages/123456`
 - gRPC: `GetMessage(message_id: "123456")`

 - HTTP: `GET /v1/users/me/messages/123456`
 - gRPC: `GetMessage(user_id: "me" message_id: "123456")`

 Rules for HTTP mapping

 1. Leaf request fields (recursive expansion nested messages in the request
    message) are classified into three categories:
    - Fields referred by the path template. They are passed via the URL path.
    - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They
    are passed via the HTTP
      request body.
    - All other fields are passed via the URL query parameters, and the
      parameter name is the field path in the request message. A repeated
      field can be represented as multiple query parameters under the same
      name.
  2. If [HttpRule.body][google.api.HttpRule.body] is "*", there is no URL
  query parameter, all fields
     are passed via URL path and HTTP request body.
  3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP
  request body, all
     fields are passed via URL path and URL query parameters.

 Path template syntax

     Template = "/" Segments [ Verb ] ;
     Segments = Segment { "/" Segment } ;
     Segment  = "*" | "**" | LITERAL | Variable ;
     Variable = "{" FieldPath [ "=" Segments ] "}" ;
     FieldPath = IDENT { "." IDENT } ;
     Verb     = ":" LITERAL ;

 The syntax `*` matches a single URL path segment. The syntax `**` matches
 zero or more URL path segments, which must be the last part of the URL path
 except the `Verb`.

 The syntax `Variable` matches part of the URL path as specified by its
 template. A variable template must not contain other variables. If a variable
 matches a single path segment, its template may be omitted, e.g. `{var}`
 is equivalent to `{var=*}`.

 The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
 contains any reserved character, such characters should be percent-encoded
 before the matching.

 If a variable contains exactly one path segment, such as `"{var}"` or
 `"{var=*}"`, when such a variable is expanded into a URL path on the client
 side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
 server side does the reverse decoding. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{var}`.

 If a variable contains multiple path segments, such as `"{var=foo/*}"`
 or `"{var=**}"`, when such a variable is expanded into a URL path on the
 client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded.
 The server side does the reverse decoding, except "%2F" and "%2f" are left
 unchanged. Such variables show up in the
 [Discovery
 Document](https://developers.google.com/discovery/v1/reference/apis) as
 `{+var}`.

 Using gRPC API Service Configuration

 gRPC API Service Configuration (service config) is a configuration language
 for configuring a gRPC service to become a user-facing product. The
 service config is simply the YAML representation of the `google.api.Service`
 proto message.

 As an alternative to annotating your proto file, you can configure gRPC
 transcoding in your service config YAML files. You do this by specifying a
 `HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
 effect as the proto annotation. This can be particularly useful if you
 have a proto that is reused in multiple services. Note that any transcoding
 specified in the service config will override any matching transcoding
 configuration in the proto.

 The following example selects a gRPC method and applies an `HttpRule` to it:

     http:
       rules:
         - selector: example.v1.Messaging.GetMessage
           get: /v1/messages/{message_id}/{sub.subfield}

 Special notes

 When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the
 proto to JSON conversion must follow the [proto3
 specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

 While the single segment variable follows the semantics of
 [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
 Expansion, the multi segment variable **does not** follow RFC 6570 Section
 3.2.3 Reserved Expansion. The reason is that the Reserved Expansion
 does not expand special characters like `?` and `#`, which would lead
 to invalid URLs. As the result, gRPC Transcoding uses a custom encoding
 for multi segment variables.

 The path variables **must not** refer to any repeated or mapped field,
 because client libraries are not capable of handling such variable expansion.

 The path variables **must not** capture the leading "/" character. The reason
 is that the most common use case "{var}" does not capture the leading "/"
 character. For consistency, all path variables must share the same behavior.

 Repeated message fields must not be mapped to URL query parameters, because
 no client library can support such complicated mapping.

 If an API needs to use a JSON array for request or response body, it can map
 the request or response body to a repeated field. However, some gRPC
 Transcoding implementations may not support this feature.


�
�
 �� Selects a method to which this rule applies.

 Refer to [selector][google.api.DocumentationRule.selector] for syntax
 details.


 �

 �	

 �
�
 ��� Determines the URL pattern is matched by this rules. This pattern can be
 used with any of the {get|put|post|delete|patch} methods. A custom method
 can be defined using the 'custom' field.


 �
\
�N Maps to HTTP GET. Used for listing and getting information about
 resources.


�


�

�
@
�2 Maps to HTTP PUT. Used for replacing a resource.


�


�

�
X
�J Maps to HTTP POST. Used for creating a resource or performing an action.


�


�

�
B
�4 Maps to HTTP DELETE. Used for deleting a resource.


�


�

�
A
�3 Maps to HTTP PATCH. Used for updating a resource.


�


�

�
�
�!� The custom pattern is used for specifying an HTTP method that is not
 included in the `pattern` field, such as HEAD, or "*" to leave the
 HTTP method unspecified for this rule. The wild-card rule is useful
 for services that provide content to Web (HTML) clients.


�

�

� 
�
�� The name of the request field whose value is mapped to the HTTP request
 body, or `*` for mapping all request fields not captured by the path
 pattern to the HTTP body, or omitted for not having any HTTP request body.

 NOTE: the referred field must be present at the top-level of the request
 message type.


�

�	

�
�
�� Optional. The name of the response field whose value is mapped to the HTTP
 response body. When omitted, the entire response message will be used
 as the HTTP response body.

 NOTE: The referred field must be present at the top-level of the response
 message type.


�

�	

�
�
	�-� Additional HTTP bindings for the selector. Nested bindings must
 not contain an `additional_bindings` field themselves (that is,
 the nesting may only be one level deep).


	�


	�

	�'

	�*,
G
� �9 A custom pattern is used for defining custom HTTP verb.


�
2
 �$ The name of this custom HTTP verb.


 �

 �	

 �
5
�' The path matched by this custom verb.


�

�	

�bproto3
��
 google/protobuf/descriptor.protogoogle.protobuf"M
FileDescriptorSet8
file (2$.google.protobuf.FileDescriptorProtoRfile"�
FileDescriptorProto
name (	Rname
package (	Rpackage

dependency (	R
dependency+
public_dependency
 (RpublicDependency'
weak_dependency (RweakDependencyC
message_type (2 .google.protobuf.DescriptorProtoRmessageTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeA
service (2'.google.protobuf.ServiceDescriptorProtoRserviceC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extension6
options (2.google.protobuf.FileOptionsRoptionsI
source_code_info	 (2.google.protobuf.SourceCodeInfoRsourceCodeInfo
syntax (	Rsyntax2
edition (2.google.protobuf.EditionRedition"�
DescriptorProto
name (	Rname;
field (2%.google.protobuf.FieldDescriptorProtoRfieldC
	extension (2%.google.protobuf.FieldDescriptorProtoR	extensionA
nested_type (2 .google.protobuf.DescriptorProtoR
nestedTypeA
	enum_type (2$.google.protobuf.EnumDescriptorProtoRenumTypeX
extension_range (2/.google.protobuf.DescriptorProto.ExtensionRangeRextensionRangeD

oneof_decl (2%.google.protobuf.OneofDescriptorProtoR	oneofDecl9
options (2.google.protobuf.MessageOptionsRoptionsU
reserved_range	 (2..google.protobuf.DescriptorProto.ReservedRangeRreservedRange#
reserved_name
 (	RreservedNamez
ExtensionRange
start (Rstart
end (Rend@
options (2&.google.protobuf.ExtensionRangeOptionsRoptions7
ReservedRange
start (Rstart
end (Rend"�
ExtensionRangeOptionsX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOptionY
declaration (22.google.protobuf.ExtensionRangeOptions.DeclarationB�Rdeclaration7
features2 (2.google.protobuf.FeatureSetRfeaturesm
verification (28.google.protobuf.ExtensionRangeOptions.VerificationState:
UNVERIFIEDB�Rverification�
Declaration
number (Rnumber
	full_name (	RfullName
type (	Rtype
reserved (Rreserved
repeated (RrepeatedJ"4
VerificationState
DECLARATION 

UNVERIFIED*	�����"�
FieldDescriptorProto
name (	Rname
number (RnumberA
label (2+.google.protobuf.FieldDescriptorProto.LabelRlabel>
type (2*.google.protobuf.FieldDescriptorProto.TypeRtype
	type_name (	RtypeName
extendee (	Rextendee#
default_value (	RdefaultValue
oneof_index	 (R
oneofIndex
	json_name
 (	RjsonName7
options (2.google.protobuf.FieldOptionsRoptions'
proto3_optional (Rproto3Optional"�
Type
TYPE_DOUBLE

TYPE_FLOAT

TYPE_INT64
TYPE_UINT64

TYPE_INT32
TYPE_FIXED64
TYPE_FIXED32
	TYPE_BOOL
TYPE_STRING	

TYPE_GROUP

TYPE_MESSAGE

TYPE_BYTES
TYPE_UINT32
	TYPE_ENUM
TYPE_SFIXED32
TYPE_SFIXED64
TYPE_SINT32
TYPE_SINT64"C
Label
LABEL_OPTIONAL
LABEL_REPEATED
LABEL_REQUIRED"c
OneofDescriptorProto
name (	Rname7
options (2.google.protobuf.OneofOptionsRoptions"�
EnumDescriptorProto
name (	Rname?
value (2).google.protobuf.EnumValueDescriptorProtoRvalue6
options (2.google.protobuf.EnumOptionsRoptions]
reserved_range (26.google.protobuf.EnumDescriptorProto.EnumReservedRangeRreservedRange#
reserved_name (	RreservedName;
EnumReservedRange
start (Rstart
end (Rend"�
EnumValueDescriptorProto
name (	Rname
number (Rnumber;
options (2!.google.protobuf.EnumValueOptionsRoptions"�
ServiceDescriptorProto
name (	Rname>
method (2&.google.protobuf.MethodDescriptorProtoRmethod9
options (2.google.protobuf.ServiceOptionsRoptions"�
MethodDescriptorProto
name (	Rname

input_type (	R	inputType
output_type (	R
outputType8
options (2.google.protobuf.MethodOptionsRoptions0
client_streaming (:falseRclientStreaming0
server_streaming (:falseRserverStreaming"�	
FileOptions!
java_package (	RjavaPackage0
java_outer_classname (	RjavaOuterClassname5
java_multiple_files
 (:falseRjavaMultipleFilesD
java_generate_equals_and_hash (BRjavaGenerateEqualsAndHash:
java_string_check_utf8 (:falseRjavaStringCheckUtf8S
optimize_for	 (2).google.protobuf.FileOptions.OptimizeMode:SPEEDRoptimizeFor

go_package (	R	goPackage5
cc_generic_services (:falseRccGenericServices9
java_generic_services (:falseRjavaGenericServices5
py_generic_services (:falseRpyGenericServices%

deprecated (:falseR
deprecated.
cc_enable_arenas (:trueRccEnableArenas*
objc_class_prefix$ (	RobjcClassPrefix)
csharp_namespace% (	RcsharpNamespace!
swift_prefix' (	RswiftPrefix(
php_class_prefix( (	RphpClassPrefix#
php_namespace) (	RphpNamespace4
php_metadata_namespace, (	RphpMetadataNamespace!
ruby_package- (	RrubyPackage7
features2 (2.google.protobuf.FeatureSetRfeaturesX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption":
OptimizeMode	
SPEED
	CODE_SIZE
LITE_RUNTIME*	�����J*+J&'Rphp_generic_services"�
MessageOptions<
message_set_wire_format (:falseRmessageSetWireFormatL
no_standard_descriptor_accessor (:falseRnoStandardDescriptorAccessor%

deprecated (:falseR
deprecated
	map_entry (RmapEntryV
&deprecated_legacy_json_field_conflicts (BR"deprecatedLegacyJsonFieldConflicts7
features (2.google.protobuf.FeatureSetRfeaturesX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	�����JJJJ	J	
"�
FieldOptionsA
ctype (2#.google.protobuf.FieldOptions.CType:STRINGRctype
packed (RpackedG
jstype (2$.google.protobuf.FieldOptions.JSType:	JS_NORMALRjstype
lazy (:falseRlazy.
unverified_lazy (:falseRunverifiedLazy%

deprecated (:falseR
deprecated
weak
 (:falseRweak(
debug_redact (:falseRdebugRedactK
	retention (2-.google.protobuf.FieldOptions.OptionRetentionR	retentionH
targets (2..google.protobuf.FieldOptions.OptionTargetTypeRtargetsW
edition_defaults (2,.google.protobuf.FieldOptions.EditionDefaultReditionDefaults7
features (2.google.protobuf.FeatureSetRfeaturesU
feature_support (2,.google.protobuf.FieldOptions.FeatureSupportRfeatureSupportX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOptionZ
EditionDefault2
edition (2.google.protobuf.EditionRedition
value (	Rvalue�
FeatureSupportG
edition_introduced (2.google.protobuf.EditionReditionIntroducedG
edition_deprecated (2.google.protobuf.EditionReditionDeprecated/
deprecation_warning (	RdeprecationWarningA
edition_removed (2.google.protobuf.EditionReditionRemoved"/
CType

STRING 
CORD
STRING_PIECE"5
JSType
	JS_NORMAL 
	JS_STRING
	JS_NUMBER"U
OptionRetention
RETENTION_UNKNOWN 
RETENTION_RUNTIME
RETENTION_SOURCE"�
OptionTargetType
TARGET_TYPE_UNKNOWN 
TARGET_TYPE_FILE
TARGET_TYPE_EXTENSION_RANGE
TARGET_TYPE_MESSAGE
TARGET_TYPE_FIELD
TARGET_TYPE_ONEOF
TARGET_TYPE_ENUM
TARGET_TYPE_ENUM_ENTRY
TARGET_TYPE_SERVICE
TARGET_TYPE_METHOD	*	�����JJ"�
OneofOptions7
features (2.google.protobuf.FeatureSetRfeaturesX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	�����"�
EnumOptions
allow_alias (R
allowAlias%

deprecated (:falseR
deprecatedV
&deprecated_legacy_json_field_conflicts (BR"deprecatedLegacyJsonFieldConflicts7
features (2.google.protobuf.FeatureSetRfeaturesX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	�����J"�
EnumValueOptions%

deprecated (:falseR
deprecated7
features (2.google.protobuf.FeatureSetRfeatures(
debug_redact (:falseRdebugRedactU
feature_support (2,.google.protobuf.FieldOptions.FeatureSupportRfeatureSupportX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	�����"�
ServiceOptions7
features" (2.google.protobuf.FeatureSetRfeatures%

deprecated! (:falseR
deprecatedX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption*	�����"�
MethodOptions%

deprecated! (:falseR
deprecatedq
idempotency_level" (2/.google.protobuf.MethodOptions.IdempotencyLevel:IDEMPOTENCY_UNKNOWNRidempotencyLevel7
features# (2.google.protobuf.FeatureSetRfeaturesX
uninterpreted_option� (2$.google.protobuf.UninterpretedOptionRuninterpretedOption"P
IdempotencyLevel
IDEMPOTENCY_UNKNOWN 
NO_SIDE_EFFECTS

IDEMPOTENT*	�����"�
UninterpretedOptionA
name (2-.google.protobuf.UninterpretedOption.NamePartRname)
identifier_value (	RidentifierValue,
positive_int_value (RpositiveIntValue,
negative_int_value (RnegativeIntValue!
double_value (RdoubleValue!
string_value (RstringValue'
aggregate_value (	RaggregateValueJ
NamePart
	name_part (	RnamePart!
is_extension (RisExtension"�


FeatureSet�
field_presence (2).google.protobuf.FeatureSet.FieldPresenceB?����EXPLICIT��IMPLICIT��EXPLICIT���RfieldPresencel
	enum_type (2$.google.protobuf.FeatureSet.EnumTypeB)����CLOSED��	OPEN���RenumType�
repeated_field_encoding (21.google.protobuf.FeatureSet.RepeatedFieldEncodingB-����EXPANDED��PACKED���RrepeatedFieldEncoding~
utf8_validation (2*.google.protobuf.FeatureSet.Utf8ValidationB)����	NONE��VERIFY���Rutf8Validation~
message_encoding (2+.google.protobuf.FeatureSet.MessageEncodingB&����LENGTH_PREFIXED���RmessageEncoding�
json_format (2&.google.protobuf.FeatureSet.JsonFormatB9�����LEGACY_BEST_EFFORT��
ALLOW���R
jsonFormat"\
FieldPresence
FIELD_PRESENCE_UNKNOWN 
EXPLICIT
IMPLICIT
LEGACY_REQUIRED"7
EnumType
ENUM_TYPE_UNKNOWN 
OPEN

CLOSED"V
RepeatedFieldEncoding#
REPEATED_FIELD_ENCODING_UNKNOWN 

PACKED
EXPANDED"I
Utf8Validation
UTF8_VALIDATION_UNKNOWN 

VERIFY
NONE""S
MessageEncoding
MESSAGE_ENCODING_UNKNOWN 
LENGTH_PREFIXED
	DELIMITED"H

JsonFormat
JSON_FORMAT_UNKNOWN 	
ALLOW
LEGACY_BEST_EFFORT*��N*�N�N*�N�NJ��"�
FeatureSetDefaultsX
defaults (2<.google.protobuf.FeatureSetDefaults.FeatureSetEditionDefaultRdefaultsA
minimum_edition (2.google.protobuf.EditionRminimumEditionA
maximum_edition (2.google.protobuf.EditionRmaximumEdition�
FeatureSetEditionDefault2
edition (2.google.protobuf.EditionReditionN
overridable_features (2.google.protobuf.FeatureSetRoverridableFeaturesB
fixed_features (2.google.protobuf.FeatureSetRfixedFeaturesJJRfeatures"�
SourceCodeInfoD
location (2(.google.protobuf.SourceCodeInfo.LocationRlocation�
Location
path (BRpath
span (BRspan)
leading_comments (	RleadingComments+
trailing_comments (	RtrailingComments:
leading_detached_comments (	RleadingDetachedComments"�
GeneratedCodeInfoM

annotation (2-.google.protobuf.GeneratedCodeInfo.AnnotationR
annotation�

Annotation
path (BRpath
source_file (	R
sourceFile
begin (Rbegin
end (RendR
semantic (26.google.protobuf.GeneratedCodeInfo.Annotation.SemanticRsemantic"(
Semantic
NONE 
SET	
ALIAS*�
Edition
EDITION_UNKNOWN 
EDITION_LEGACY�
EDITION_PROTO2�
EDITION_PROTO3�
EDITION_2023�
EDITION_2024�
EDITION_1_TEST_ONLY
EDITION_2_TEST_ONLY
EDITION_99997_TEST_ONLY��
EDITION_99998_TEST_ONLY��
EDITION_99999_TEST_ONLY��
EDITION_MAX����B~
com.google.protobufBDescriptorProtosHZ-google.golang.org/protobuf/types/descriptorpb��GPB�Google.Protobuf.ReflectionJ��
& �

�
& 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
2� Author: kenton@google.com (Kenton Varda)
  Based on original Protocol Buffers design by
  Sanjay Ghemawat, Jeff Dean, and others.

 The messages in this file describe the definitions found in .proto files.
 A valid .proto file can be translated directly to a FileDescriptorProto
 without any other information (e.g. without reading its imports).


( 

* D
	
* D

+ ,
	
+ ,

, 1
	
, 1

- 7
	
%- 7

. !
	
$. !

/ 
	
/ 

3 

	3 t descriptor.proto must be optimized for speed because reflection-based
 algorithms don't work during bootstrapping.

j
 7 9^ The protocol compiler can output a FileDescriptorSet containing the .proto
 files it parses.



 7

  8(

  8


  8

  8#

  8&'
-
 < ]! The full set of known editions.



 <
:
  >- A placeholder for an unknown edition value.


  >

  >
�
 B� A placeholder edition for specifying default behaviors *before* a feature
 was first introduced.  This is effectively an "infinite past".


 B

 B
�
 H� Legacy syntax "editions".  These pre-date editions, but behave much like
 distinct editions.  These can't be used to specify the edition of proto
 files, but feature definitions must supply proto2/proto3 defaults for
 backwards compatibility.


 H

 H

 I

 I

 I
�
 N� Editions that have been released.  The specific values are arbitrary and
 should not be depended on, but they will always be time-ordered for easy
 comparison.


 N

 N

 O

 O

 O
}
 Sp Placeholder editions for testing feature resolution.  These should not be
 used or relyed on outside of tests.


 S

 S

 T

 T

 T

 U"

 U

 U!

 	V"

 	V

 	V!

 
W"

 
W

 
W!
�
 \� Placeholder for specifying unbounded edition support.  This should only
 ever be used by plugins that can expect to never require any changes to
 support a new edition.


 \

 \
0
` �# Describes a complete .proto file.



`
9
 a", file name, relative to root of source tree


 a


 a

 a

 a
*
b" e.g. "foo", "foo.bar", etc.


b


b

b

b
4
e!' Names of files imported by this file.


e


e

e

e 
Q
g(D Indexes of the public imported files in the dependency list above.


g


g

g"

g%'
z
j&m Indexes of the weak imported files in the dependency list.
 For Google-internal migration only. Do not use.


j


j

j 

j#%
6
m,) All top-level definitions in this file.


m


m

m'

m*+

n-

n


n

n(

n+,

o.

o


o!

o")

o,-

p.

p


p

p )

p,-

	r#

	r


	r

	r

	r!"
�

x/� This field contains optional information about the original source code.
 You may safely remove this entire field without harming runtime
 functionality of the descriptors -- the information is needed only by
 development tools.



x



x


x*


x-.
�
~� The syntax of the proto file.
 The supported values are "proto2", "proto3", and "editions".

 If `edition` is present, this value must be "editions".


~


~

~

~
.
�   The edition of the proto file.


�


�

�

�
)
� � Describes a message type.


�

 �

 �


 �

 �

 �

�*

�


�

� %

�()

�.

�


�

� )

�,-

�+

�


�

�&

�)*

�-

�


�

�(

�+,

 ��

 �


  �" Inclusive.


  �

  �

  �

  �

 �" Exclusive.


 �

 �

 �

 �

 �/

 �

 �"

 �#*

 �-.

�.

�


�

�)

�,-

�/

�


�

� *

�-.

�&

�


�

�!

�$%
�
��� Range of reserved tag numbers. Reserved tag numbers may not be used by
 fields or extension ranges in the same message. Reserved ranges may
 not overlap.


�


 �" Inclusive.


 �

 �

 �

 �

�" Exclusive.


�

�

�

�

�,

�


�

�'

�*+
�
	�%u Reserved field names, which may not be used by fields in the same message.
 A given name may only be reserved once.


	�


	�

	�

	�"$

� �

�
O
 �:A The parser stores options it doesn't recognize here. See above.


 �


 �

 �3

 �69

 ��

 �

K
  �; The extension number declared within the extension range.


  �

  �

  �

  �
z
 �"j The fully-qualified name of the extension field. There must be a leading
 dot in front of the full name.


 �

 �

 �

 � !
�
 �� The fully-qualified type name of the extension field. Unlike
 Metadata.type, Declaration.type must have a leading dot for messages
 and enums.


 �

 �

 �

 �
�
 �� If true, indicates that the number is reserved in the extension range,
 and any extension field with the number will fail to compile. Set this
 when a declared extension field is deleted.


 �

 �

 �

 �
�
 �z If true, indicates that the extension must be defined as repeated.
 Otherwise the extension must be defined as optional.


 �

 �

 �

 �
$
 	�" removed is_repeated


 	 �

 	 �

 	 �
�
�F� For external users: DO NOT USE. We are in the process of open sourcing
 extension declaration and executing internal cleanups before it can be
 used externally.


�


�

�"

�%&

�'E

�(D
=
�$/ Any features defined in the specific edition.


�


�

�

�!#
@
 ��0 The verification state of the extension range.


 �
C
  �3 All the extensions of the range must be declared.


  �

  �

 �

 �

 �
�
��;~ The verification state of the range.
 TODO: flip the default to DECLARATION once all empty ranges
 are marked as UNVERIFIED.


�


�

�)

�,-

�:

�

�9
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �
3
� �% Describes a field within a message.


�

 ��

 �
S
  �C 0 is reserved for errors.
 Order is weird for historical reasons.


  �

  �

 �

 �

 �
w
 �g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
 negative values are likely.


 �

 �

 �

 �

 �
w
 �g Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
 negative values are likely.


 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �
�
 	�� Tag-delimited aggregate.
 Group type is deprecated and not supported after google.protobuf. However, Proto3
 implementations should still be able to parse the group wire format and
 treat group fields as unknown fields.  In Editions, the group wire format
 can be enabled via the `message_encoding` feature.


 	�

 	�
-
 
�" Length-delimited aggregate.


 
�

 
�
#
 � New in version 2.


 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 �
'
 �" Uses ZigZag encoding.


 �

 �
'
 �" Uses ZigZag encoding.


 �

 �

��

�
*
 � 0 is reserved for errors


 �

 �

�

�

�
�
�� The required label is only allowed in google.protobuf.  In proto3 and Editions
 it's explicitly prohibited.  In Editions, the `field_presence` feature
 can be used to get this behavior.


�

�

 �

 �


 �

 �

 �

�

�


�

�

�

�

�


�

�

�
�
�� If type_name is set, this need not be set.  If both this and type_name
 are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.


�


�

�

�
�
� � For message and enum types, this is the name of the type.  If the name
 starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
 rules are used to find the type (i.e. first the nested types within this
 message are searched, then within the parent, on up to the root
 namespace).


�


�

�

�
~
�p For extensions, this is the name of the type being extended.  It is
 resolved in the same manner as type_name.


�


�

�

�
�
�$� For numeric types, contains the original text representation of the value.
 For booleans, "true" or "false".
 For strings, contains the default text contents (not escaped in any way).
 For bytes, contains the C escaped value.  All bytes >= 128 are escaped.


�


�

�

�"#
�
�!v If set, gives the index of a oneof in the containing type's oneof_decl
 list.  This field is a member of that oneof.


�


�

�

� 
�
�!� JSON name of this field. The value is set by protocol compiler. If the
 user has set a "json_name" option on this field, that option's value
 will be used. Otherwise, it's deduced from the field's name by converting
 it to camelCase.


�


�

�

� 

	�$

	�


	�

	�

	�"#
�	

�%�	 If true, this is a proto3 "optional". When a proto3 field is optional, it
 tracks presence regardless of field type.

 When proto3_optional is true, this field must belong to a oneof to signal
 to old proto3 clients that presence is tracked for this field. This oneof
 is known as a "synthetic" oneof, and this field must be its sole member
 (each proto3 optional field gets its own synthetic oneof). Synthetic oneofs
 exist in the descriptor only, and do not generate any API. Synthetic oneofs
 must be ordered after all "real" oneofs.

 For message fields, proto3_optional doesn't create any semantic change,
 since non-repeated message fields always track presence. However it still
 indicates the semantic detail of whether the user wrote "optional" or not.
 This can be useful for round-tripping the .proto file. For consistency we
 give message fields a synthetic oneof also, even though it is not required
 to track presence. This is especially important because the parser can't
 tell if a field is a message or an enum, so it must always create a
 synthetic oneof.

 Proto2 optional fields do not set this flag, because they already indicate
 optional with `LABEL_OPTIONAL`.



�



�


�


�"$
"
� � Describes a oneof.


�

 �

 �


 �

 �

 �

�$

�


�

�

�"#
'
� � Describes an enum type.


�

 �

 �


 �

 �

 �

�.

�


�#

�$)

�,-

�#

�


�

�

�!"
�
 ��� Range of reserved numeric values. Reserved values may not be used by
 entries in the same enum. Reserved ranges may not overlap.

 Note that this is distinct from DescriptorProto.ReservedRange in that it
 is inclusive such that it can appropriately represent the entire int32
 domain.


 �


  �" Inclusive.


  �

  �

  �

  �

 �" Inclusive.


 �

 �

 �

 �
�
�0� Range of reserved numeric values. Reserved numeric values may not be used
 by enum values in the same enum declaration. Reserved ranges may not
 overlap.


�


�

�+

�./
l
�$^ Reserved enum value names, which may not be reused. A given name may only
 be reserved once.


�


�

�

�"#
1
� �# Describes a value within an enum.


� 

 �

 �


 �

 �

 �

�

�


�

�

�

�(

�


�

�#

�&'
$
� � Describes a service.


�

 �

 �


 �

 �

 �

�,

�


� 

�!'

�*+

�&

�


�

�!

�$%
0
	� �" Describes a method of a service.


	�

	 �

	 �


	 �

	 �

	 �
�
	�!� Input and output type names.  These are resolved in the same way as
 FieldDescriptorProto.type_name, but must refer to a message type.


	�


	�

	�

	� 

	�"

	�


	�

	�

	� !

	�%

	�


	�

	� 

	�#$
E
	�77 Identifies if client streams multiple client messages


	�


	�

	� 

	�#$

	�%6

	�05
E
	�77 Identifies if server streams multiple server messages


	�


	�

	� 

	�#$

	�%6

	�05
�

� �2N ===================================================================
 Options
2� Each of the definitions above may have "options" attached.  These are
 just annotations which may cause code to be generated slightly differently
 or may contain hints for code that manipulates protocol messages.

 Clients may define custom options as extensions of the *Options messages.
 These extensions may not yet be known at parsing time, so the parser cannot
 store the values in them.  Instead it stores them in a field in the *Options
 message called uninterpreted_option. This field must have the same name
 across all *Options messages. We then use this field to populate the
 extensions when we build a descriptor, at which point all protos have been
 parsed and so all extensions are known.

 Extension numbers for custom options may be chosen as follows:
 * For options which will only be used within a single application or
   organization, or for experimental options, use field numbers 50000
   through 99999.  It is up to you to ensure that you do not use the
   same number for multiple options.
 * For options which will be published and used publicly by multiple
   independent entities, e-mail protobuf-global-extension-registry@google.com
   to reserve extension numbers. Simply provide your project name (e.g.
   Objective-C plugin) and your project website (if available) -- there's no
   need to explain how you intend to use them. Usually you only need one
   extension number. You can declare multiple options with only one extension
   number by putting them in a sub-message. See the Custom Options section of
   the docs for examples:
   https://developers.google.com/protocol-buffers/docs/proto#options
   If this turns out to be popular, a web service will be set up
   to automatically assign option numbers.



�
�

 �#� Sets the Java package where classes generated from this .proto will be
 placed.  By default, the proto package is used, but this is often
 inappropriate because proto packages do not normally start with backwards
 domain names.



 �



 �


 �


 �!"
�

�+� Controls the name of the wrapper Java class generated for the .proto file.
 That class will always contain the .proto file's getDescriptor() method as
 well as any top-level extensions defined in the .proto file.
 If java_multiple_files is disabled, then all the other classes from the
 .proto file will be nested inside the single wrapper outer class.



�



�


�&


�)*
�

�;� If enabled, then the Java code generator will generate a separate .java
 file for each top-level message, enum, and service defined in the .proto
 file.  Thus, these types will *not* be nested inside the wrapper class
 named by java_outer_classname.  However, the wrapper class will still be
 generated to contain the file's getDescriptor() method as well as any
 top-level extensions defined in the file.



�



�


�#


�&(


�):


�49
)

�E This option does nothing.



�



�


�-


�02


�3D


�4C
�

�>� A proto2 file can set this to true to opt in to UTF-8 checking for Java,
 which will throw an exception if invalid UTF-8 is parsed from the wire or
 assigned to a string field.

 TODO: clarify exactly what kinds of field types this option
 applies to, and update these docs accordingly.

 Proto3 files already perform these checks. Setting the option explicitly to
 false has no effect: it cannot be used to opt proto3 files out of UTF-8
 checks.



�



�


�&


�)+


�,=


�7<
L

 ��< Generated classes can be optimized for speed or code size.



 �
D

  �"4 Generate complete code for parsing, serialization,



  �	


  �
G

 � etc.
"/ Use ReflectionOps to implement these methods.



 �


 �
G

 �"7 Generate code using MessageLite and the lite runtime.



 �


 �


�;


�



�


�$


�'(


�):


�49
�

�"� Sets the Go package where structs generated from this .proto will be
 placed. If omitted, the Go package will be derived from the following:
   - The basename of the package import path, if provided.
   - Otherwise, the package statement in the .proto file, if present.
   - Otherwise, the basename of the .proto file, without extension.



�



�


�


�!
�

�;� Should generic services be generated in each language?  "Generic" services
 are not specific to any particular RPC system.  They are generated by the
 main code generators in each language (without additional plugins).
 Generic services were the only kind of service generation supported by
 early versions of google.protobuf.

 Generic services are now considered deprecated in favor of using plugins
 that generate code specific to your particular RPC system.  Therefore,
 these default to false.  Old code which depends on generic services should
 explicitly set them to true.



�



�


�#


�&(


�):


�49


�=


�



�


�%


�(*


�+<


�6;


	�;


	�



	�


	�#


	�&(


	�):


	�49
+

	�" removed php_generic_services



	 �


	 �


	 �



�"



 �!
�


�2� Is this file deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for everything in the file, or it will be completely ignored; in the very
 least, this is a formalization for deprecating files.




�




�



�



�



� 1



�+0


�7q Enables the use of arenas for the proto messages in this file. This applies
 only to generated classes for C++.



�



�


� 


�#%


�&6


�15
�

�)� Sets the objective c class prefix which is prepended to all objective c
 generated classes from this .proto. There is no default.



�



�


�#


�&(
I

�(; Namespace for generated classes; defaults to the package.



�



�


�"


�%'
�

�$� By default Swift generators will take the proto package and CamelCase it
 replacing '.' with underscore and use that to prefix the types/symbols
 defined. When this options is provided, they will use this value instead
 to prefix the types/symbols defined.



�



�


�


�!#
~

�(p Sets the php class prefix which is prepended to all php generated classes
 from this .proto. Default is empty.



�



�


�"


�%'
�

�%� Use this option to change the namespace of php generated classes. Default
 is empty. When this option is empty, the package name will be used for
 determining the namespace.



�



�


�


�"$
�

�.� Use this option to change the namespace of php generated metadata classes.
 Default is empty. When this option is empty, the proto file name will be
 used for determining the namespace.



�



�


�(


�+-
�

�$� Use this option to change the package of ruby generated classes. Default
 is empty. When this option is not set, the package name will be used for
 determining the ruby package.



�



�


�


�!#
=

�$/ Any features defined in the specific edition.



�



�


�


�!#
|

�:n The parser stores options it doesn't recognize here.
 See the documentation for the "Options" section above.



�



�


�3


�69
�

�z Clients can define custom options in extensions of this message.
 See the documentation for the "Options" section above.



 �


 �


 �


	�


	�


	�


	�

� �

�
�
 �>� Set true to use the old proto1 MessageSet wire format for extensions.
 This is provided for backwards-compatibility with the MessageSet wire
 format.  You should not use this for any other reason:  It's less
 efficient, has fewer features, and is more complicated.

 The message must be defined exactly as follows:
   message Foo {
     option message_set_wire_format = true;
     extensions 4 to max;
   }
 Note that the message cannot have any defined fields; MessageSets only
 have extensions.

 All extensions of your type must be singular messages; e.g. they cannot
 be int32s, enums, or repeated messages.

 Because this is an option, the above two restrictions are not enforced by
 the protocol compiler.


 �


 �

 �'

 �*+

 �,=

 �7<
�
�F� Disables the generation of the standard "descriptor()" accessor, which can
 conflict with a field of the same name.  This is meant to make migration
 from proto1 easier; new code should avoid fields named "descriptor".


�


�

�/

�23

�4E

�?D
�
�1� Is this message deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the message, or it will be completely ignored; in the very least,
 this is a formalization for deprecating messages.


�


�

�

�

�0

�*/

	�

	 �

	 �

	 �

	�

	�

	�

	�

	�

	�
�
�� Whether the message is an automatically generated map entry type for the
 maps field.

 For maps fields:
     map<KeyType, ValueType> map_field = 1;
 The parsed descriptor looks like:
     message MapFieldEntry {
         option map_entry = true;
         optional KeyType key = 1;
         optional ValueType value = 2;
     }
     repeated MapFieldEntry map_field = 1;

 Implementations may choose not to generate the map_entry=true message, but
 use a native map in the target language to hold the keys and values.
 The reflection APIs in such implementations still need to work as
 if the field is a repeated message field.

 NOTE: Do not set the option in .proto files. Always use the maps syntax
 instead. The option should only be implicitly set by the proto compiler
 parser.


�


�

�

�
$
	�" javalite_serializable


	�

	�

	�

	�" javanano_as_lite


	�

	�

	�
�
�P� Enable the legacy handling of JSON field name conflicts.  This lowercases
 and strips underscored from the fields before comparison in proto3 only.
 The new behavior takes `json_name` into account and applies to proto2 as
 well.

 This should only be used as a temporary measure against broken builds due
 to the change in behavior for JSON field name conflicts.

 TODO This is legacy behavior we plan to remove once downstream
 teams have had time to migrate.


�


�

�6

�9;

�<O

�=N
=
�$/ Any features defined in the specific edition.


�


�

�

�!#
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

� �

�
�
 �E� NOTE: ctype is deprecated. Use `features.(pb.cpp).string_type` instead.
 The ctype option instructs the C++ code generator to use a different
 representation of the field than it normally would.  See the specific
 options below.  This option is only implemented to support use of
 [ctype=CORD] and [ctype=STRING] (the default) on non-repeated fields of
 type "bytes" in the open source release.
 TODO: make ctype actually deprecated.


 �


 �

 �

 �

 �D

 �=C

 ��

 �

  � Default mode.


  �


  �
�
 �� The option [ctype=CORD] may be applied to a non-repeated field of type
 "bytes". It indicates that in C++, the data should be stored in a Cord
 instead of a string.  For very large strings, this may reduce memory
 fragmentation. It may also allow better performance when parsing from a
 Cord, or when parsing with aliasing enabled, as the parsed Cord may then
 alias the original buffer.


 �

 �

 �

 �

 �
�
�� The packed option can be enabled for repeated primitive fields to enable
 a more efficient representation on the wire. Rather than repeatedly
 writing the tag and type for each element, the entire array is encoded as
 a single length-delimited blob. In proto3, only explicit setting it to
 false will avoid using packed encoding.  This option is prohibited in
 Editions, but the `repeated_field_encoding` feature can be used to control
 the behavior.


�


�

�

�
�
�3� The jstype option determines the JavaScript type used for values of the
 field.  The option is permitted only for 64 bit integral and fixed types
 (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
 is represented as JavaScript string, which avoids loss of precision that
 can happen when a large value is converted to a floating point JavaScript.
 Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
 use the JavaScript "number" type.  The behavior of the default option
 JS_NORMAL is implementation dependent.

 This option is an enum to permit additional types to be added, e.g.
 goog.math.Integer.


�


�

�

�

�2

�(1

��

�
'
 � Use the default type.


 �

 �
)
� Use JavaScript strings.


�

�
)
� Use JavaScript numbers.


�

�
�

�+�
 Should this field be parsed lazily?  Lazy applies only to message-type
 fields.  It means that when the outer message is initially parsed, the
 inner message's contents will not be parsed but instead stored in encoded
 form.  The inner message will actually be parsed when it is first accessed.

 This is only a hint.  Implementations are free to choose whether to use
 eager or lazy parsing regardless of the value of this option.  However,
 setting this option true suggests that the protocol author believes that
 using lazy parsing on this field is worth the additional bookkeeping
 overhead typically needed to implement it.

 This option does not affect the public interface of any generated code;
 all method signatures remain the same.  Furthermore, thread-safety of the
 interface is not affected by this option; const methods remain safe to
 call from multiple threads concurrently, while non-const methods continue
 to require exclusive access.

 Note that lazy message fields are still eagerly verified to check
 ill-formed wireformat or missing required fields. Calling IsInitialized()
 on the outer message would fail if the inner message has missing required
 fields. Failed verification would result in parsing failure (except when
 uninitialized messages are acceptable).


�


�

�

�

�*

�$)
�
�7� unverified_lazy does no correctness checks on the byte stream. This should
 only be used where lazy with verification is prohibitive for performance
 reasons.


�


�

�

�"$

�%6

�05
�
�1� Is this field deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for accessors, or it will be completely ignored; in the very least, this
 is a formalization for deprecating fields.


�


�

�

�

�0

�*/
?
�,1 For Google-internal migration only. Do not use.


�


�

�

�

�+

�%*
�
�4� Indicate that the field value should not be printed out when using debug
 formats, e.g. when the field contains sensitive credentials.


�


�

�

�!

�"3

�-2
�
��� If set to RETENTION_SOURCE, the option will be omitted from the binary.
 Note: as of January 2023, support for this is in progress and does not yet
 have an effect (b/264593489).


�

 �

 �

 �

�

�

�

�

�

�

�*

�


�

�$

�')
�
��� This indicates the types of entities that the field may apply to when used
 as an option. If it is unset, then the field may be freely used as an
 option on any kind of entity. Note: as of January 2023, support for this is
 in progress and does not yet have an effect (b/264593489).


�

 �

 �

 �

�

�

�

�$

�

�"#

�

�

�

�

�

�

�

�

�

�

�

�

�

�

�

�

�

�

	�

	�

	�

	�)

	�


	�

	�#

	�&(

 ��

 �


  �!

  �

  �

  �

  � 
"
 �" Textproto value.


 �

 �

 �

 �


�0


�



�


�*


�-/
=
�$/ Any features defined in the specific edition.


�


�

�

�!#
D
��4 Information about the support window of a feature.


�

�
 �,� The edition that this feature was first available in.  In editions
 earlier than this one, the default assigned to EDITION_LEGACY will be
 used, and proto files will not be able to override it.


 �

 �

 �'

 �*+
w
�,g The edition this feature becomes deprecated in.  Using this after this
 edition may trigger warnings.


�

�

�'

�*+
v
�,f The deprecation warning text if this feature is used after the edition it
 was marked deprecated in.


�

�

�'

�*+
�
�)� The edition this feature is no longer available in.  In editions after
 this one, the last default assigned will be used, and proto files will
 not be able to override it.


�

�

�$

�'(

�/

�


�

�)

�,.
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

	�" removed jtype


	 �

	 �

	 �
9
	�", reserve target, target_obsolete_do_not_use


	�

	�

	�

� �

�
=
 �#/ Any features defined in the specific edition.


 �


 �

 �

 �!"
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

� �

�
`
 � R Set this option to true to allow mapping different tag names to the same
 value.


 �


 �

 �

 �
�
�1� Is this enum deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum, or it will be completely ignored; in the very least, this
 is a formalization for deprecating enums.


�


�

�

�

�0

�*/

	�" javanano_as_lite


	 �

	 �

	 �
�
�O� Enable the legacy handling of JSON field name conflicts.  This lowercases
 and strips underscored from the fields before comparison in proto3 only.
 The new behavior takes `json_name` into account and applies to proto2 as
 well.
 TODO Remove this legacy behavior once downstream teams have
 had time to migrate.


�


�

�6

�9:

�;N

�<M
=
�#/ Any features defined in the specific edition.


�


�

�

�!"
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

� �

�
�
 �1� Is this enum value deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the enum value, or it will be completely ignored; in the very least,
 this is a formalization for deprecating enum values.


 �


 �

 �

 �

 �0

 �*/
=
�#/ Any features defined in the specific edition.


�


�

�

�!"
�
�3� Indicate that fields annotated with this enum value should not be printed
 out when using debug formats, e.g. when the field contains sensitive
 credentials.


�


�

�

� 

�!2

�,1
H
�;: Information about the support window of a feature value.


�


�&

�'6

�9:
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

� �

�
=
 �$/ Any features defined in the specific edition.


 �


 �

 �

 �!#
�
�2� Is this service deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the service, or it will be completely ignored; in the very least,
 this is a formalization for deprecating services.
2� Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


�


�

�

�

� 1

�+0
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �

� �

�
�
 �2� Is this method deprecated?
 Depending on the target platform, this can emit Deprecated annotations
 for the method, or it will be completely ignored; in the very least,
 this is a formalization for deprecating methods.
2� Note:  Field numbers 1 through 32 are reserved for Google's internal RPC
   framework.  We apologize for hoarding these numbers to ourselves, but
   we were already using them long before we decided to release Protocol
   Buffers.


 �


 �

 �

 �

 � 1

 �+0
�
 ��� Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
 or neither? HTTP based RPC implementation may choose GET verb for safe
 methods, and PUT verb for idempotent methods instead of the default POST.


 �

  �

  �

  �
$
 �" implies idempotent


 �

 �
7
 �"' idempotent, but may have side effects


 �

 �

��&

�


�

�-

�02

�%

�$
=
�$/ Any features defined in the specific edition.


�


�

�

�!#
O
�:A The parser stores options it doesn't recognize here. See above.


�


�

�3

�69
Z
�M Clients can define custom options in extensions of this message. See above.


 �

 �

 �
�
� �� A message representing a option the parser does not recognize. This only
 appears in options protos created by the compiler::Parser class.
 DescriptorPool resolves these when building Descriptor objects. Therefore,
 options protos in descriptor objects (e.g. returned by Descriptor::options(),
 or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
 in them.


�
�
 ��� The name of the uninterpreted option.  Each string represents a segment in
 a dot-separated name.  is_extension is true iff a segment represents an
 extension (denoted with parentheses in options specs in .proto files).
 E.g.,{ ["foo", false], ["bar.baz", true], ["moo", false] } represents
 "foo.(bar.baz).moo".


 �


  �"

  �

  �

  �

  � !

 �#

 �

 �

 �

 �!"

 �

 �


 �

 �

 �
�
�'� The value of the uninterpreted option, in whatever type the tokenizer
 identified it as during parsing. Exactly one of these should be set.


�


�

�"

�%&

�)

�


�

�$

�'(

�(

�


�

�#

�&'

�#

�


�

�

�!"

�"

�


�

�

� !

�&

�


�

�!

�$%
�
� �� TODO Enums in C++ gencode (and potentially other languages) are
 not well scoped.  This means that each of the feature enums below can clash
 with each other.  The short names we've chosen maximize call-site
 readability, but leave us very open to this scenario.  A future feature will
 be designed and implemented to handle this, hopefully before we ever hit a
 conflict here.
2O ===================================================================
 Features


�

 ��

 �

  �

  �

  �

 �

 �

 �

 �

 �

 �

 �

 �

 �

 ��

 �


 �

 �'

 �*+

 �,�

 �!

  �

 �

 ��

  �E

 �E

 �C

��

�

 �

 �

 �

�

�

�

�

�


�

��

�


�

�

� !

�"�

�!

 �

�

��

 �C

�A

��

�

 �(

 �#

 �&'

�

�


�

�

�

�

��

�


� 

�!8

�;<

�=�

�!

 �

�

��

 �E

�C

��

�

 � 

 �

 �

�

�


�

�

�

�

�

 �

 �

 �

��

�


�

�)

�,-

�.�

�!

 �

�

��

 �A

�C

��

�

 �!

 �

 � 

�

�

�

�

�

�

��

�


�

�+

�./

�0�

�!

 �

�

��

 �L

��

�

 �

 �

 �

�

�	

�

�

�

�

��

�


�

�!

�$%

�&�

�!

 �!

�

�

��

 �O

�B

	�

	 �

	 �

	 �

��

 �

 �

 �
#
�" For internal testing


�

�

�
:
�"- for https://github.com/bufbuild/protobuf-es


�

�

�
�
� �� A compiled specification for the defaults of a set of features.  These
 messages are generated from FeatureSet extensions and can be used to seed
 feature resolution. The resolution with this object becomes a simple search
 for the closest matching edition, followed by proto merges.


�
�
 ��� A map from every known edition with a unique set of defaults to its
 defaults. Not all editions may be contained here.  For a given edition,
 the defaults at the closest matching edition ordered at or before it should
 be used.  This field must be in strict ascending order by edition.


 �
"

  �!

  �

  �

  �

  � 
N
 �1> Defaults of features that can be overridden in this edition.


 �

 �

 �,

 �/0
P
 �+@ Defaults of features that can't be overridden in this edition.


 �

 �

 �&

 �)*

 	�

 	 �

 	 �

 	 �

 	�

 	�

 	�

 
�

 
 �

 �1

 �


 �#

 �$,

 �/0
�
�'t The minimum supported edition (inclusive) when this was constructed.
 Editions before this will not have defaults.


�


�

�"

�%&
�
�'x The maximum known edition (inclusive) when this was constructed. Editions
 after this will not have reliable defaults.


�


�

�"

�%&
�
� �	j Encapsulates information about the original source file from which a
 FileDescriptorProto was generated.
2` ===================================================================
 Optional source code info


�
�
 �	!� A Location identifies a piece of source code in a .proto file which
 corresponds to a particular definition.  This information is intended
 to be useful to IDEs, code indexers, documentation generators, and similar
 tools.

 For example, say we have a file like:
   message Foo {
     optional string foo = 1;
   }
 Let's look at just the field definition:
   optional string foo = 1;
   ^       ^^     ^^  ^  ^^^
   a       bc     de  f  ghi
 We have the following locations:
   span   path               represents
   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).

 Notes:
 - A location may refer to a repeated field itself (i.e. not to any
   particular index within it).  This is used whenever a set of elements are
   logically enclosed in a single code segment.  For example, an entire
   extend block (possibly containing multiple extension definitions) will
   have an outer location whose path refers to the "extensions" repeated
   field without an index.
 - Multiple locations may have the same path.  This happens when a single
   logical declaration is spread out across multiple places.  The most
   obvious example is the "extend" block again -- there may be multiple
   extend blocks in the same scope, each of which will have the same path.
 - A location's span is not always a subset of its parent's span.  For
   example, the "extendee" of an extension declaration appears at the
   beginning of the "extend" block and is shared by all extensions within
   the block.
 - Just because a location's span is a subset of some other location's span
   does not mean that it is a descendant.  For example, a "group" defines
   both a type and a field in a single declaration.  Thus, the locations
   corresponding to the type and field and their components will overlap.
 - Code which tries to interpret locations should probably be designed to
   ignore those that it doesn't understand, as more types of locations could
   be recorded in the future.


 �	


 �	

 �	

 �	 

 �	�	

 �	

�
  �	,� Identifies which part of the FileDescriptorProto was defined at this
 location.

 Each element is a field number or an index.  They form a path from
 the root FileDescriptorProto to the place where the definition appears.
 For example, this path:
   [ 4, 3, 2, 7, 1 ]
 refers to:
   file.message_type(3)  // 4, 3
       .field(7)         // 2, 7
       .name()           // 1
 This is because FileDescriptorProto.message_type has field number 4:
   repeated DescriptorProto message_type = 4;
 and DescriptorProto.field has field number 2:
   repeated FieldDescriptorProto field = 2;
 and FieldDescriptorProto.name has field number 1:
   optional string name = 1;

 Thus, the above path gives the location of a field name.  If we removed
 the last element:
   [ 4, 3, 2, 7 ]
 this path refers to the whole field declaration (from the beginning
 of the label to the terminating semicolon).


  �	

  �	

  �	

  �	

  �	+

  �	*
�
 �	,� Always has exactly three or four elements: start line, start column,
 end line (optional, otherwise assumed same as start line), end column.
 These are packed into a single field for efficiency.  Note that line
 and column numbers are zero-based -- typically you will want to add
 1 to each before displaying to a user.


 �	

 �	

 �	

 �	

 �	+

 �	*
�
 �	)� If this SourceCodeInfo represents a complete declaration, these are any
 comments appearing before and after the declaration which appear to be
 attached to the declaration.

 A series of line comments appearing on consecutive lines, with no other
 tokens appearing on those lines, will be treated as a single comment.

 leading_detached_comments will keep paragraphs of comments that appear
 before (but not connected to) the current element. Each paragraph,
 separated by empty lines, will be one comment element in the repeated
 field.

 Only the comment content is provided; comment markers (e.g. //) are
 stripped out.  For block comments, leading whitespace and an asterisk
 will be stripped from the beginning of each line other than the first.
 Newlines are included in the output.

 Examples:

   optional int32 foo = 1;  // Comment attached to foo.
   // Comment attached to bar.
   optional int32 bar = 2;

   optional string baz = 3;
   // Comment attached to baz.
   // Another line attached to baz.

   // Comment attached to moo.
   //
   // Another line attached to moo.
   optional double moo = 4;

   // Detached comment for corge. This is not leading or trailing comments
   // to moo or corge because there are blank lines separating it from
   // both.

   // Detached comment for corge paragraph 2.

   optional string corge = 5;
   /* Block comment attached
    * to corge.  Leading asterisks
    * will be removed. */
   /* Block comment attached to
    * grault. */
   optional int32 grault = 6;

   // ignored detached comments.


 �	

 �	

 �	$

 �	'(

 �	*

 �	

 �	

 �	%

 �	()

 �	2

 �	

 �	

 �	-

 �	01
�
�	 �
� Describes the relationship between generated code and its original source
 file. A GeneratedCodeInfo message is associated with only one generated
 source file, but may contain references to different source .proto files.


�	
x
 �	%j An Annotation connects some span of text in generated code to an element
 of its generating .proto file.


 �	


 �	

 �	 

 �	#$

 �	�


 �	

�
  �	, Identifies the element in the original source .proto file. This field
 is formatted the same as SourceCodeInfo.Location.path.


  �	

  �	

  �	

  �	

  �	+

  �	*
O
 �	$? Identifies the filesystem path to the original source .proto.


 �	

 �	

 �	

 �	"#
w
 �	g Identifies the starting offset in bytes in the generated code
 that relates to the identified object.


 �	

 �	

 �	

 �	
�
 �
� Identifies the ending offset in bytes in the generated code that
 relates to the identified object. The end offset should be one past
 the last relevant byte (so the length of the text = end - begin).


 �


 �


 �


 �

j
  �
�
X Represents the identified object's effect on the element in the original
 .proto file.


  �
	
F
   �
4 There is no effect or the effect is indescribable.


	   �



	   �

<
  �
* The element is set or otherwise mutated.


	  �
	

	  �

8
  �
& An alias to the element is returned.


	  �


	  �


 �
#

 �


 �


 �


 �
!"
�	
google/api/annotations.proto
google.apigoogle/api/http.proto google/protobuf/descriptor.proto:K
http.google.protobuf.MethodOptions�ʼ" (2.google.api.HttpRuleRhttpB�
com.google.apiBAnnotationsProtoPZpgithub.com/Yux77Yux/platform_backend/generated/google.golang.org/genproto/googleapis/api/annotations;annotations�GAPIJ�
 
�
 2� Copyright 2024 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.


 
	
  
	
 *
	
 �


 �

 "
	

 "

 1
	
 1

 '
	
 '

 "
	
$ "
	
 

  See `HttpRule`.



 $


 



 


 bproto3
�1
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB�
com.google.protobufBTimestampProtoPZ2google.golang.org/protobuf/types/known/timestamppb��GPB�Google.Protobuf.WellKnownTypesJ�/
 �
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" 
	
" 

# I
	
# I

$ ,
	
$ ,

% /
	
% /

& "
	

& "

' !
	
$' !

( ;
	
%( ;
�
 � �� A Timestamp represents a point in time independent of any time zone or local
 calendar, encoded as a count of seconds and fractions of seconds at
 nanosecond resolution. The count is relative to an epoch at UTC midnight on
 January 1, 1970, in the proleptic Gregorian calendar which extends the
 Gregorian calendar backwards to year one.

 All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
 second table is needed for interpretation, using a [24-hour linear
 smear](https://developers.google.com/time/smear).

 The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
 restricting to that range, we ensure that we can convert to and from [RFC
 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

 # Examples

 Example 1: Compute Timestamp from POSIX `time()`.

     Timestamp timestamp;
     timestamp.set_seconds(time(NULL));
     timestamp.set_nanos(0);

 Example 2: Compute Timestamp from POSIX `gettimeofday()`.

     struct timeval tv;
     gettimeofday(&tv, NULL);

     Timestamp timestamp;
     timestamp.set_seconds(tv.tv_sec);
     timestamp.set_nanos(tv.tv_usec * 1000);

 Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

     FILETIME ft;
     GetSystemTimeAsFileTime(&ft);
     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
     Timestamp timestamp;
     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

 Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

     long millis = System.currentTimeMillis();

     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
         .setNanos((int) ((millis % 1000) * 1000000)).build();

 Example 5: Compute Timestamp from Java `Instant.now()`.

     Instant now = Instant.now();

     Timestamp timestamp =
         Timestamp.newBuilder().setSeconds(now.getEpochSecond())
             .setNanos(now.getNano()).build();

 Example 6: Compute Timestamp from current time in Python.

     timestamp = Timestamp()
     timestamp.GetCurrentTime()

 # JSON Mapping

 In JSON format, the Timestamp type is encoded as a string in the
 [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 where {year} is always expressed using four digits while {month}, {day},
 {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 is required. A proto3 JSON serializer should always use UTC (as indicated by
 "Z") when printing the Timestamp type and a proto3 JSON parser should be
 able to accept both UTC and other timezones (as indicated by an offset).

 For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 01:30 UTC on January 15, 2017.

 In JavaScript, one can convert a Date object to this format using the
 standard
 [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
 method. In Python, a standard `datetime.datetime` object can be converted
 to this format using
 [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
 the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
 the Joda Time's [`ISODateTimeFormat.dateTime()`](
 http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
 ) to obtain a formatter capable of generating timestamps in this format.



 �
�
  �� Represents seconds of UTC time since Unix epoch
 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
 9999-12-31T23:59:59Z inclusive.


  �

  �

  �
�
 �� Non-negative fractions of a second at nanosecond resolution. Negative
 second values with fractions must still have non-negative nanos values
 that count forward in time. Must be from 0 to 999,999,999
 inclusive.


 �

 �

 �bproto3
�
common/access_token.protocommongoogle/protobuf/timestamp.proto"^
AccessToken
value (	Rvalue9

expires_at (2.google.protobuf.TimestampR	expiresAtB7Z5github.com/Yux77Yux/platform_backend/generated/commonJ�
  

  

 

 L
	
 L
$
  )" 引入 Timestamp 类型



  


 

  	

  	

  		

  	
!
 
+" Token 到期时间


 


 
&

 
)*bproto3
�
common/after_auth.protocommon"d
	AfterAuth
user_id (RuserId
creation_id (R
creationId

comment_id (R	commentId"G
AnyAfterAuth7
any_after_auth (2.common.AfterAuthRanyAfterAuthB7Z5github.com/Yux77Yux/platform_backend/generated/commonJ�
  

  

 

 L
	
 L


  



 

  

  

  

  

 

 

 

 

 	

 	

 	

 	


 




 (

 


 

 #

 &'bproto3
�
common/api_response.protocommon"�
ApiResponse2
status (2.common.ApiResponse.StatusRstatus
code (	Rcode
message (	Rmessage
details (	Rdetails
trace_id (	RtraceId"9
Status
SUCCESS 	
ERROR
PENDING

FAILEDB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M


  


 

  

  	

   " 操作成功


   

   

  	" 操作失败


  	

  	
&
  
" 操作正在处理中


  


  

#
  " 操作彻底失败


  

  
7
  "* 操作状态（例如: success, error）


  


  

  
3
 "& 状态码（例如: 200, 400, 500）


 


 

 
$
 " 用户友好的消息


 


 

 
*
 " 错误详情或附加信息


 


 

 
*
 " 请求追踪 ID（可选）


 


 

 bproto3
�
common/creation.common.protocommon"

CreationId
id (Rid"2
ViewCreation
id (Rid
ipv4 (	Ripv4B7Z5github.com/Yux77Yux/platform_backend/generated/commonJ�
  

  

 

 L
	
 L


  


 

  

  

  


  



 





 

 

 


 





	

bproto3
�
common/custom_options.protocommon google/protobuf/descriptor.proto:^
min_user_credentials_length.google.protobuf.FieldOptionsц (RminUserCredentialsLength:^
max_user_credentials_length.google.protobuf.FieldOptions҆ (RmaxUserCredentialsLength:P
max_user_name_length.google.protobuf.FieldOptionsۆ (RmaxUserNameLength:N
max_user_bio_length.google.protobuf.FieldOptions܆ (RmaxUserBioLength:I
max_title_length.google.protobuf.FieldOptions� (RmaxTitleLength:W
max_introduction_length.google.protobuf.FieldOptions� (RmaxIntroductionLengthB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M
	
  *
	
 
	
 	.


 #


 		


 	
%


 	(-
	
.


#


	



%


(-
	
'


#


	






!&
	
&


#


	






 %
	
#


#


	






"
	
*


#


	



!


$)bproto3
�
common/operate.protocommon*g
Operate
NONE 
VIEW
LIKE
CANCEL_COLLECT
COLLECT
CANCEL_LIKE
DEL_VIEWB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M


  


 

  

  

  	


 

 

 	


 	

 	

 		


 


 


 


 

 	

 

 

 

 

 

 


 bproto3
�.
google/protobuf/any.protogoogle.protobuf"6
Any
type_url (	RtypeUrl
value (RvalueBv
com.google.protobufBAnyProtoPZ,google.golang.org/protobuf/types/known/anypb�GPB�Google.Protobuf.WellKnownTypesJ�,
 �
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" C
	
" C

# ,
	
# ,

$ )
	
$ )

% "
	

% "

& !
	
$& !

' ;
	
%' ;
�
  �� `Any` contains an arbitrary serialized protocol buffer message along with a
 URL that describes the type of the serialized message.

 Protobuf library provides support to pack/unpack Any values in the form
 of utility functions or additional generated methods of the Any type.

 Example 1: Pack and unpack a message in C++.

     Foo foo = ...;
     Any any;
     any.PackFrom(foo);
     ...
     if (any.UnpackTo(&foo)) {
       ...
     }

 Example 2: Pack and unpack a message in Java.

     Foo foo = ...;
     Any any = Any.pack(foo);
     ...
     if (any.is(Foo.class)) {
       foo = any.unpack(Foo.class);
     }
     // or ...
     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
       foo = any.unpack(Foo.getDefaultInstance());
     }

  Example 3: Pack and unpack a message in Python.

     foo = Foo(...)
     any = Any()
     any.Pack(foo)
     ...
     if any.Is(Foo.DESCRIPTOR):
       any.Unpack(foo)
       ...

  Example 4: Pack and unpack a message in Go

      foo := &pb.Foo{...}
      any, err := anypb.New(foo)
      if err != nil {
        ...
      }
      ...
      foo := &pb.Foo{}
      if err := any.UnmarshalTo(foo); err != nil {
        ...
      }

 The pack methods provided by protobuf library will by default use
 'type.googleapis.com/full.type.name' as the type URL and the unpack
 methods only use the fully qualified type name after the last '/'
 in the type URL, for example "foo.bar.com/x/y.z" will yield type
 name "y.z".

 JSON
 ====
 The JSON representation of an `Any` value uses the regular
 representation of the deserialized, embedded message, with an
 additional field `@type` which contains the type URL. Example:

     package google.profile;
     message Person {
       string first_name = 1;
       string last_name = 2;
     }

     {
       "@type": "type.googleapis.com/google.profile.Person",
       "firstName": <string>,
       "lastName": <string>
     }

 If the embedded message type is well-known and has a custom JSON
 representation, that representation will be embedded adding a field
 `value` which holds the custom JSON in addition to the `@type`
 field. Example (for message [google.protobuf.Duration][]):

     {
       "@type": "type.googleapis.com/google.protobuf.Duration",
       "value": "1.212s"
     }




 
�
  �� A URL/resource name that uniquely identifies the type of the serialized
 protocol buffer message. This string must contain at least
 one "/" character. The last segment of the URL's path must represent
 the fully qualified name of the type (as in
 `path/google.protobuf.Duration`). The name should be in a canonical form
 (e.g., leading "." is not accepted).

 In practice, teams usually precompile into the binary all types that they
 expect it to use in the context of Any. However, for URLs which use the
 scheme `http`, `https`, or no scheme, one can optionally set up a type
 server that maps type URLs to message definitions as follows:

 * If no scheme is provided, `https` is assumed.
 * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   value in binary format, or produce an error.
 * Applications are allowed to cache lookup results based on the
   URL, or have them precompiled into a binary to avoid any
   lookup. Therefore, binary compatibility needs to be preserved
   on changes to types. (Use versioned type names to manage
   breaking changes.)

 Note: this functionality is not currently available in the official
 protobuf release, and it is not used for type URLs beginning with
 type.googleapis.com. As of May 2023, there are no widely used type server
 implementations and no plans to implement one.

 Schemes other than `http`, `https` (or the empty scheme) might be
 used with implementation specific semantics.



  �

  �	

  �
W
 �I Must be a valid serialized protocol buffer of the above specified type.


 �

 �

 �bproto3
�
common/trace-wrap.protocommongoogle/protobuf/any.proto"q
Wrapper
trace_id (	RtraceId
	full_name (	RfullName.
payload (2.google.protobuf.AnyRpayloadB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M
	
  #


  


 

  	

  	

  		

  	

 


 


 
	

 


 "

 

 

  !bproto3
�
common/update_count.protocommoncommon/creation.common.protocommon/operate.proto"[

UserAction"
id (2.common.CreationIdRid)
operate (2.common.OperateRoperate"=
AnyUserAction,
actions (2.common.UserActionRactionsB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M
	
  &
	
 


 
 


 


  

  

  

  

 

 

 

 


 




 "

 


 

 

  !bproto3
�
common/user_default.protocommoncommon/custom_options.proto"I
UserDefault
user_id (RuserId!
	user_name (	BصdRuserNameB8Z6github.com/Yux77Yux/platform_backend/generated/common;J�
  

  

 

 M
	
 M
	
  %


  


 

  	

  		

  	


  	

 
?

 



 


 


 
>

 ۆ
=bproto3
�
"common/user_creation_comment.protocommoncommon/user_default.protocommon/custom_options.proto"�
UserCreationComment6
user_default (2.common.UserDefaultRuserDefault
user_avatar (	R
userAvatar 
user_bio (	B��RuserBio
	followers (R	followersB7Z5github.com/Yux77Yux/platform_backend/generated/commonJ�
  

  

 

 L
	
 L
	
  #
	
 %


 	 


 	

  
(

  


  
#

  
&'

 

 


 

 

 >

 


 

 

 =

 ܆<

 

 	

 


 bproto3
�
!event/aggregator/aggregator.protoevent.aggregator*Q
Exchange
EXCHANGE_INCREASE_VIEW )
%EXCHANGE_UPDATE_CREATION_ACTION_COUNT*H
Queue
QUEUE_INCREASE_VIEW &
"QUEUE_UPDATE_CREATION_ACTION_COUNT*I

RoutingKey
KEY_INCREASE_VIEW $
 KEY_UPDATE_CREATION_ACTION_COUNTBBZ@github.com/Yux77Yux/platform_backend/generated/event/aggregator;J�
  

  

 

 W
	
 W


  	


 

  

  

  

 ,

 '

 *+


 





 

 

 

)

$

'(


 




 

 

 

'

"

%&bproto3
�
event/user/user.proto
event.user*�
Exchange
EXCHANGE_REGISTER 
EXCHANGE_STORE_USER
EXCHANGE_STORE_CREDENTIAL
EXCHANGE_UPDATE_USER_SPACE
EXCHANGE_UPDATE_USER_BIO
EXCHANGE_UPDATE_USER_AVATAR
EXCHANGE_FOLLOW
EXCHANGE_CANCEL_FOLLOW
EXCHANGE_UPDATE_USER_STATUS
EXCHANGE_DEL_REVIEWER	*�
Queue
QUEUE_REGISTER 
QUEUE_STORE_USER
QUEUE_STORE_CREDENTIAL
QUEUE_UPDATE_USER_SPACE
QUEUE_UPDATE_USER_BIO
QUEUE_UPDATE_USER_AVATAR
QUEUE_FOLLOW
QUEUE_CANCEL_FOLLOW
QUEUE_UPDATE_USER_STATUS
QUEUE_DEL_REVIEWER	*�

RoutingKey
KEY_REGISTER 
KEY_STORE_USER
KEY_STORE_CREDENTIAL
KEY_UPDATE_USER_SPACE
KEY_UPDATE_USER_BIO
KEY_UPDATE_USER_AVATAR

KEY_FOLLOW
KEY_CANCEL_FOLLOW
KEY_UPDATE_USER_STATUS
KEY_DEL_REVIEWER	B<Z:github.com/Yux77Yux/platform_backend/generated/event/user;J�

  .

  

 

 Q
	
 Q


  


 

  

  

  

 

 

 

 	 

 	

 	

 
!

 


 
 

 

 

 

 "

 

  !

 

 

 

 

 

 

 "

 

  !

 	

 	

 	


  





 

 

 

















































	

	

	


" .


"

 #

 #

 #

$

$

$

%

%

%

&

&

&

'

'

'

(

(

(

)

)

)

*

*

*

+

+

+

	,

	,

	,bproto3
�
event/creation/creation.protoevent.creation*�
Exchange
EXCHANGE_UPDATE_DB_CREATION 
EXCHANGE_STORE_CREATION"
EXCHANGE_UPDATE_CACHE_CREATION#
EXCHANGE_UPDATE_CREATION_STATUS
EXCHANGE_DELETE_CREATION)
%EXCHANGE_UPDATE_CREATION_ACTION_COUNT
EXCHANGE_PEND_CREATION*�
Queue
QUEUE_UPDATE_DB_CREATION 
QUEUE_STORE_CREATION
QUEUE_UPDATE_CACHE_CREATION 
QUEUE_UPDATE_CREATION_STATUS
QUEUE_DELETE_CREATION&
"QUEUE_UPDATE_CREATION_ACTION_COUNT*�

RoutingKey
KEY_UPDATE_DB_CREATION 
KEY_STORE_CREATION
KEY_UPDATE_CACHE_CREATION
KEY_UPDATE_CREATION_STATUS
KEY_DELETE_CREATION$
 KEY_UPDATE_CREATION_ACTION_COUNT
KEY_PEND_CREATIONB@Z>github.com/Yux77Yux/platform_backend/generated/event/creation;J�
  !

  

 

 U
	
 U


  


 

  "

  

   !

 

 

 

 	%

 	 

 	#$

 
&

 
!

 
$%

 

 

 

 ,

 '

 *+

 

 

 


 





 

 

 







"



 !

#



!"







)

$

'(


 !




 

 

 







 





!



 







'

"

%&

 

 

 bproto3
�
#event/interaction/interaction.protoevent.interaction*�
Exchange
EXCHANGE_COMPUTE_CREATION 
EXCHANGE_COMPUTE_USER
EXCHANGE_UPDATE_DB
EXCHANGE_BATCH_UPDATE_DB
EXCHANGE_ADD_COLLECTION
EXCHANGE_ADD_LIKE
EXCHANGE_ADD_VIEW
EXCHANGE_CANCEL_LIKE)
%EXCHANGE_UPDATE_CREATION_ACTION_COUNT*�
Queue
QUEUE_COMPUTE_CREATION 
QUEUE_COMPUTE_USER
QUEUE_UPDATE_DB
QUEUE_BATCH_UPDATE_DB
QUEUE_ADD_COLLECTION
QUEUE_ADD_LIKE
QUEUE_ADD_VIEW
QUEUE_CANCEL_LIKE*�

RoutingKey
KEY_COMPUTE_CREATION 
KEY_COMPUTE_USER
KEY_UPDATE_DB
KEY_BATCH_UPDATE_DB
KEY_ADD_COLLECTION
KEY_ADD_LIKE
KEY_ADD_VIEW
KEY_CANCEL_LIKE$
 KEY_UPDATE_CREATION_ACTION_COUNTBCZAgithub.com/Yux77Yux/platform_backend/generated/event/interaction;J�	
  '

  

 

 X
	
 X


  


 

   

  

  

 

 

 

 	

 	

 	

 


 


 


 

 

 

 

 

 

 

 

 

 

 

 

 ,

 '

 *+


 





 

 

 












































 '




 

 

 







 

 

 

!

!

!

"

"

"

#

#

#

$

$

$

%

%

%

&'

&"

&%&bproto3
�
event/review/review.protoevent.review*�
Exchange
EXCHANGE_COMMENT_REVIEW 
EXCHANGE_USER_REVIEW
EXCHANGE_CREATION_REVIEW
EXCHANGE_NEW_REVIEW
EXCHANGE_UPDATE
EXCHANGE_BATCH_UPDATE
EXCHANGE_PEND_CREATION
EXCHANGE_UPDATE_USER_STATUS#
EXCHANGE_UPDATE_CREATION_STATUS
EXCHANGE_DELETE_CREATION	
EXCHANGE_DELETE_COMMENT
*�
Queue
QUEUE_COMMENT_REVIEW 
QUEUE_USER_REVIEW
QUEUE_CREATION_REVIEW
QUEUE_NEW_REVIEW
QUEUE_UPDATE
QUEUE_BATCH_UPDATE
QUEUE_PEND_CREATION*�

RoutingKey
KEY_COMMENT_REVIEW 
KEY_USER_REVIEW
KEY_CREATION_REVIEW
KEY_NEW_REVIEW

KEY_UPDATE
KEY_BATCH_UPDATE
KEY_PEND_CREATION
KEY_UPDATE_USER_STATUS
KEY_UPDATE_CREATION_STATUS
KEY_DELETE_CREATION	
KEY_DELETE_COMMENT
B>Z<github.com/Yux77Yux/platform_backend/generated/event/review;J�

  *

  

 

 S
	
 S


  


 

  

  

  

 

 

 

 	

 	

 	

 


 


 


 

 

 

 

 

 

 

 

 

 "

 

  !

 &

 !

 $%

 	

 	

 	

 


 


 



 





 

 

 






































 *




 

 

 

 

 

 

!

!

!

"

"

"

#

#

#

$

$

$

%

%

%

&

&

&

'!

'

' 

	(

	(

	(


)


)


)bproto3
�
event/comment/comment.protoevent.comment*E
Exchange
EXCHANGE_PUBLISH_COMMENT 
EXCHANGE_DELETE_COMMENT*<
Queue
QUEUE_PUBLISH_COMMENT 
QUEUE_DELETE_COMMENT*=

RoutingKey
KEY_PUBLISH_COMMENT 
KEY_DELETE_COMMENTB?Z=github.com/Yux77Yux/platform_backend/generated/event/comment;J�
  

  

 

 T
	
 T


  	


 

  

  

  

 

 

 


 





 

 

 








 




 

 

 





bproto3
�
'creation/messages/creation_status.protocreation.messages*Q
CreationStatus	
DRAFT 
PENDING
	PUBLISHED
REJECTED

DELETEB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N


  


 

  

  

  


 

 	

 

 	

 	

 	

 


 



 


 

 

 bproto3
�
'creation/messages/creation_upload.protocreation.messages'creation/messages/creation_status.proto"�
CreationUpload
	author_id (RauthorId
src (	Rsrc
	thumbnail (	R	thumbnail
title (	Rtitle
bio (	Rbio9
status (2!.creation.messages.CreationStatusRstatus
duration (Rduration
category_id (R
categoryIdB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N
	
  1


  


 

  	

  	

  	

  	

 


 


 
	

 


 

 

 	

 

 

 

 	

 

 

 

 	

 

 .

 "

 #)

 ,-

 

 

 

 

 

 

 

 bproto3
�
 creation/messages/creation.protocreation.messagesgoogle/protobuf/timestamp.proto'creation/messages/creation_upload.proto"�
CreationInfo7
creation (2.creation.messages.CreationRcreationV
creation_engagement (2%.creation.messages.CreationEngagementRcreationEngagement7
category (2.creation.messages.CategoryRcategory"�
Creation
creation_id (R
creationId>
	base_info (2!.creation.messages.CreationUploadRbaseInfo;
upload_time (2.google.protobuf.TimestampR
uploadTime"�
CreationEngagement
creation_id (R
creationId
views (Rviews
likes (Rlikes
saves (Rsaves=
publish_time (2.google.protobuf.TimestampRpublishTime"y
Category
category_id (R
categoryId
parent (Rparent
name (	Rname 
description (	Rdescription"M
AnyCreation>
any_creation (2.creation.messages.CreationRanyCreation"[
AnyCreationEngagementB
	any_count (2%.creation.messages.CreationEngagementRanyCountB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�	
  +

  

 

 N
	
 N
	
  )
	
 1


 
 


 


  

  


  

  

 -

 

 (

 +,

 

 


 

 


 




 

 

 

 

1

"

#,

/0

,



'

*+


 




 

 

 

 

























-



(

+,


 #




 

 

 

 

 

 

 

 

!

!

!	

!

"

"

"	

"


% '


%

 &%

 &


 &

 & 

 &#$


) +


)

 *,

 *


 *

 *'

 **+bproto3
�
'aggregator/messages/creation_card.protoaggregator.messages creation/messages/creation.protocommon/user_default.protogoogle/protobuf/timestamp.proto"�
CreationCard7
creation (2.creation.messages.CreationRcreationV
creation_engagement (2%.creation.messages.CreationEngagementRcreationEngagement'
user (2.common.UserDefaultRuser3
time_at (2.google.protobuf.TimestampRtimeAtB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  

  

 

 P
	
 P
	
  *
	
 #
	
	 )


  


 

  *

  

  %

  ()

 ?

 &

 ':

 =>

 

 

 

 

 (

 

 #

 &'bproto3
�
aggregator/methods/get.protoaggregator.methods'aggregator/messages/creation_card.protocommon/access_token.protocommon/api_response.proto"\
HistoryRequest6
access_token (2.common.AccessTokenRaccessToken
page (Rpage"E
HomeRequest6
access_token (2.common.AccessTokenRaccessToken"`
CollectionsRequest6
access_token (2.common.AccessTokenRaccessToken
page (Rpage":
SimilarCreationsRequest
creation_id (R
creationId"r
GetCardsResponse%
msg (2.common.ApiResponseRmsg7
cards (2!.aggregator.messages.CreationCardRcards"B
SearchCreationsRequest
title (	Rtitle
page (Rpage"�
SearchCreationsResponse%
msg (2.common.ApiResponseRmsg
count (Rcount7
cards (2!.aggregator.messages.CreationCardRcardsB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  +

  

 

 P
	
 P
	
  1
	
 #
	
	 #


  


 

  &

  

  !

  $%

 

 

 

 


 




 &

 

 !

 $%


 




 &

 

 !

 $%










 




 

 

 

 


  




 

 

 

 

6




+

,1

45


" %


"

 #

 #

 #	

 #

$

$

$

$


' +


'

 (

 (

 (

 (

)

)

)

)

*6

*


*+

*,1

*45bproto3
�
review/messages/type.protoreview.messages*1

TargetType
CREATION 
USER
COMMENTB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  


  

 

 L
	
 L


  



 

  

  


  

 

 

 	


 	

 		

 	bproto3
�
 review/messages/new_review.protoreview.messagesreview/messages/type.protogoogle/protobuf/timestamp.proto"�
	NewReview
id (Rid
	target_id (RtargetId<
target_type (2.review.messages.TargetTypeR
targetType9

created_at (2.google.protobuf.TimestampR	createdAt
msg (	RmsgB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L
	
  $
	
 )


 	 


 	

  
" 审核信息id


  


  



  


 " 审核目标id


 

 

 

 -" 对象类型 


 

 (

 +,

 +

 

 &

 )*
'
 " 举报信息（可选）


 

 	

 bproto3
�
review/messages/status.protoreview.messages*D
ReviewStatus
PENDING 
APPROVED
REJECTED
DELETEDB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L


  


 

  

  	

  

 

 


 

 	

 	


 	

 


 
	

 
bproto3
�
review/messages/review.protoreview.messages review/messages/new_review.protoreview/messages/status.protogoogle/protobuf/timestamp.proto"�
Review,
new (2.review.messages.NewReviewRnew
reviewer_id (R
reviewerId5
status (2.review.messages.ReviewStatusRstatus
remark (	Rremark9

updated_at (2.google.protobuf.TimestampR	updatedAt">
	AnyReview1
reviews (2.review.messages.ReviewRreviewsB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L
	
  *
	
 &
	
 )


  


 

  $

  

  

  "#

 

 

 

 

 *" 对象类型 


 

 %

 ()

 

 

 	

 

 +

 

 &

 )*


 




 

 


 

 

 bproto3
�
comment/messages/comment.protocomment.messagesgoogle/protobuf/timestamp.proto"�
Comment

comment_id (R	commentId
root (Rroot
parent (Rparent
dialog (Rdialog
user_id (RuserId
creation_id (R
creationId9

created_at (2.google.protobuf.TimestampR	createdAt
content (	Rcontent
media	 (	Rmedia"^

TopComment3
comment (2.comment.messages.CommentRcomment
	sub_count (RsubCount"h
SecondComment3
comment (2.comment.messages.CommentRcomment"
reply_user_id (RreplyUserId"�
CommentArea
creation_id (R
creationId%
total_comments (RtotalCommentsE
area_status (2$.comment.messages.CommentArea.StatusR
areaStatus"-
Status
DEFAULT 

HIDING

CLOSED"H

AnyComment:
any_comment (2.comment.messages.CommentR
anyComment"Y
AnyCommentAreaG
any_comment_area (2.comment.messages.CommentAreaRanyCommentAreaB8Z6github.com/Yux77Yux/platform_backend/generated/commentJ�

  0

  

 

 M
	
 M
	
  )


  


 

  	

  	

  	

  	

 


 


 


 


 

 

 

 

 

 

 

 

 

 

 

 

 

 

 

 

 +

 

 &

 )*

 

 

 	

 

 

 

 	

 


 




 

 	

 


 










 




 

 	

 


 










 (




  $

  

  !

  !

  !

 "

 "


 "

 #

 #


 #

 %

 %

 %

 %

&

&

&

&

'

'

'	

'


* ,


*

 +#

 +


 +

 +

 +!"


. 0


.

 /,

 /


 /

 /'

 /*+bproto3
�
aggregator/methods/review.protoaggregator.methodsreview/messages/review.protoreview/messages/status.proto creation/messages/creation.protocomment/messages/comment.proto"common/user_creation_comment.protocommon/access_token.protocommon/api_response.proto"�
GetReviewsRequest6
access_token (2.common.AccessTokenRaccessToken5
status (2.review.messages.ReviewStatusRstatus
page (Rpage"N
GetNewReviewsRequest6
access_token (2.common.AccessTokenRaccessToken"�
GetCreationReviewsResponse%
msg (2.common.ApiResponseRmsgW
reviews (2=.aggregator.methods.GetCreationReviewsResponse.CreationReviewRreviews
count (Rcountz
CreationReview/
review (2.review.messages.ReviewRreview7
creation (2.creation.messages.CreationRcreation"�
GetUserReviewsResponse%
msg (2.common.ApiResponseRmsgO
reviews (25.aggregator.methods.GetUserReviewsResponse.UserReviewRreviews
count (Rcountn

UserReview/
review (2.review.messages.ReviewRreview/
user (2.common.UserCreationCommentRuser"�
GetCommentReviewsResponse%
msg (2.common.ApiResponseRmsgU
reviews (2;.aggregator.methods.GetCommentReviewsResponse.CommentReviewRreviews
count (Rcountu
CommentReview/
review (2.review.messages.ReviewRreview3
comment (2.comment.messages.CommentRcommentB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  5

  

 

 P
	
 P
	
  &
	
 &
	
 *
	
	 (
	
 ,
	
 #
	
 #


  


 

  &

  

  !

  $%
9
 *", 获取需要审核的，或已被审核的


 

 %

 ()

 

 

 

 


 




 &

 

 !

 $%


 !


"

 

 


  &

  

  !

  $%

 ,

 

 '

 *+

 

 

 

 

&






!

$%

 

 

 

 


# +


#

 $'

 $


  %&

  %

  %!

  %$%

 &(

 &

 &#

 &&'

 (

 (

 (

 (

)"

)


)

)

) !

*

*

*

*


- 5


-!

 .1

 .


  /&

  /

  /!

  /$%

 0)

 0

 0$

 0'(

 2

 2

 2

 2

3%

3


3

3 

3#$

4

4

4

4bproto3
�
user/messages/user_role.protouser.messages*0
UserRole
USER 	
ADMIN
SUPER_ADMINB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  


  

 

 J
	
 J


  



 

  

  

  	


 

 

 


 	

 	

 	bproto3
�
user/messages/user_login.protouser.messagescommon/user_default.protouser/messages/user_role.proto"�
	UserLogin6
user_default (2.common.UserDefaultRuserDefault
user_avatar (	R
userAvatar4
	user_role (2.user.messages.UserRoleRuserRoleB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J
	
  #
	
 '


 	 


 	

  
(

  


  
#

  
&'

 

 


 

 

 )

 

 $

 '(bproto3
�
$user/messages/user_credentials.protouser.messagesuser/messages/user_role.protocommon/custom_options.proto"�
UserCredentials$
username (	B����2Rusername%
password (	B	�����Rpassword

user_email (	R	userEmail4
	user_role (2.user.messages.UserRoleRuserRole
user_id (RuserId"]
AnyUserCredentialsG
any_credentials (2.user.messages.UserCredentialsRanyCredentialsB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J
	
  '
	
 %


 	 


 	

  


  


  
	

  


  


  ц,

  ҆-

 

 

 	

 

 

 ц,

 ҆.

 

 

 	

 

 '

 

 "

 %&

 

 

 

 


 




 /

 


 

 *

 -.bproto3
�
!auth/messages/refresh_token.protoauth.messagesgoogle/protobuf/timestamp.proto"_
RefreshToken
value (	Rvalue9

expires_at (2.google.protobuf.TimestampR	expiresAtB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
$
  )" 引入 Timestamp 类型



  


 

  	

  	

  		

  	
!
 
+" Token 到期时间


 


 
&

 
)*bproto3
�
auth/messages/tokens.protoauth.messagescommon/access_token.proto!auth/messages/refresh_token.proto"�
Tokens@
refresh_token (2.auth.messages.RefreshTokenRrefreshToken6
access_token (2.common.AccessTokenRaccessTokenB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
	
  #
	
 +


 	 


 	

  
/

  


  
*

  
-.

 &

 

 !

 $%bproto3
�
aggregator/methods/login.protoaggregator.methodsuser/messages/user_login.proto$user/messages/user_credentials.protoauth/messages/tokens.protocommon/api_response.proto"Y
LoginRequestI
user_credentials (2.user.messages.UserCredentialsRuserCredentials"�
LoginResponse7

user_login (2.user.messages.UserLoginR	userLogin-
tokens (2.auth.messages.TokensRtokens%
msg (2.common.ApiResponseRmsgB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  

  

 

 P
	
 P
	
  (
	
 .
	
 $
	
	 #


  


 

  7

  !

  "2

  56


 




 +

 

 &

 )*

$





"#







bproto3
�
comment/methods/get.protocomment.methodscomment/messages/comment.protocommon/access_token.protocommon/api_response.proto"&
GetCommentsRequest
ids (Rids"9
InitialCommentsRequest
creation_id (R
creationId"�
InitialCommentsResponse8
comments (2.comment.messages.TopCommentRcomments@
comment_area (2.comment.messages.CommentAreaRcommentArea

page_count (R	pageCount%
msg (2.common.ApiResponseRmsg"L
GetTopCommentsRequest
creation_id (R
creationId
page (Rpage"y
GetTopCommentsResponse8
comments (2.comment.messages.TopCommentRcomments%
msg (2.common.ApiResponseRmsg"c
GetSecondCommentsRequest
creation_id (R
creationId
root (Rroot
page (Rpage"
GetSecondCommentsResponse;
comments (2.comment.messages.SecondCommentRcomments%
msg (2.common.ApiResponseRmsg"e
GetReplyCommentsRequest6
access_token (2.common.AccessTokenRaccessToken
page (Rpage"s
GetCommentsResponse5
comments (2.comment.messages.CommentRcomments%
msg (2.common.ApiResponseRmsg"#
GetCommentRequest
id (Rid"p
GetCommentResponse3
comment (2.comment.messages.CommentRcomment%
msg (2.common.ApiResponseRmsgB8Z6github.com/Yux77Yux/platform_backend/generated/commentJ�
  E

  

 

 M
	
 M
	
  (
	
 #
	
	 #


  


 

  

  


  

  

  

  初始化





 

 

 

 


 




 4

 


 &

 '/

 23

0



+

./
















#
  作品的一级评论





 

 

 

 










! $


!

 "4

 "


 "&

 "'/

 "23

#

#

#

#

' + 二级评论



' 

 (

 (

 (

 (
!
)" 一级评论所在


)

)

)

*

*

*

*


- 0


-!

 .7

 .


 .)

 .*2

 .56

/

/

/

/
_
4 7S 消息中心，没做
 回复我的评论，在页面的消息内显示,权限类



4

 5&" 自己的id，


 5

 5!

 5$%

6

6

6

6


9 <


9

 :1

 :


 :#

 :$,

 :/0

;

;

;

;


	> @


	>

	 ?

	 ?

	 ?


	 ?



B E



B


 C'


 C


 C"


 C%&


D


D


D


Dbproto3
�	
&aggregator/messages/comment_info.protoaggregator.messagescomment/messages/comment.proto"common/user_creation_comment.proto"�
CommentInfo>
comment_user (2.common.UserCreationCommentRcommentUser3
comment (2.comment.messages.CommentRcomment"�
TopCommentInfo>
comment_user (2.common.UserCreationCommentRcommentUser=
top_comment (2.comment.messages.TopCommentR
topComment"�
SecondCommentInfoG
second_comment (2 .aggregator.messages.CommentInfoRsecondComment:

reply_user (2.common.UserCreationCommentR	replyUserB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  

  

 

 P
	
 P
	
  (
	
 ,


 	 


 	
'
  
." 评论中的用户信息


  


  
)

  
,-

 '

 

 "

 %&


 



'
 ." 评论中的用户信息


 

 )

 ,-

.



)

,-


 



'
 !" 评论中的用户信息


 

 

  
$
," 回复对象的信息




'

*+bproto3
�
aggregator/methods/watch.protoaggregator.methods creation/messages/creation.protocomment/messages/comment.protocomment/methods/get.proto"common/user_creation_comment.protocommon/api_response.protocommon/access_token.proto&aggregator/messages/comment_info.proto"7
WatchCreationRequest
creation_id (R
creationId"�
WatchCreationResponse%
msg (2.common.ApiResponseRmsg@
creation_user (2.common.UserCreationCommentRcreationUserD
creation_info (2.creation.messages.CreationInfoRcreationInfo"[
InitialCommentsRequestA
request (2'.comment.methods.InitialCommentsRequestRrequest"�
InitialCommentsResponse?
comments (2#.aggregator.messages.TopCommentInfoRcomments1
area (2.comment.messages.CommentAreaRarea

page_count (R	pageCount%
msg (2.common.ApiResponseRmsg"Y
GetTopCommentsRequest@
request (2&.comment.methods.GetTopCommentsRequestRrequest"�
GetTopCommentsResponse?
comments (2#.aggregator.messages.TopCommentInfoRcomments%
msg (2.common.ApiResponseRmsg"�
GetSecondCommentsRequestC
request (2).comment.methods.GetSecondCommentsRequestRrequest6
access_token (2.common.AccessTokenRaccessToken"�
GetSecondCommentsResponseB
comments (2&.aggregator.messages.SecondCommentInfoRcomments%
msg (2.common.ApiResponseRmsg"�
GetReplyCommentsRequestB
request (2(.comment.methods.GetReplyCommentsRequestRrequest6
access_token (2.common.AccessTokenRaccessToken"z
GetCommentsResponse<
comments (2 .aggregator.messages.CommentInfoRcomments%
msg (2.common.ApiResponseRmsgB;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  J

  

 

 P
	
 P
	
  *
	
 (
	
 #
	

 ,
	
 #
	
 #
	
 0
\
  P 视频和用户信息   专注于返回视频信息，用户数据处理延后



 

  

  

  

  


 




 " 返回的响应


 

 

 
$
/" 作品的用户信息




*

-.
$
3" 作品的详细信息


 

!.

12
�
# %  第一次加载作品的评论
20 相似视频 Request 与 Response 在get.proto
2/ 在加载完视频之后，首先加载评论



#

 $5

 $(

 $)0

 $34


& +


&

 ';

 '


 '-

 '.6

 '9:

((

(

(#

(&'

)

)

)

)

*" 返回的响应


*

*

*


- /


-

 .4

 .'

 .(/

 .23


1 4


1

 2;

 2


 2-

 2.6

 29:

3" 返回的响应


3

3

3
&
7 : 一级评论内的评论



7 

 87

 8*

 8+2

 856

9&

9

9!

9$%


< ?


<!

 =>

 =


 =0

 =19

 =<=

>" 返回的响应


>

>

>
P
B ED 回复我的评论，在页面的消息内显示, 权限类 没做



B

 C6

 C)

 C*1

 C45

D&

D

D!

D$%


	G J


	G

	 H8

	 H


	 H*

	 H+3

	 H67

	I" 返回的响应


	I

	I

	Ibproto3
�!
aggregator/agg_service.proto
aggregatorgoogle/api/annotations.protoaggregator/methods/get.protoaggregator/methods/review.protoaggregator/methods/login.protoaggregator/methods/watch.proto2�
AggregatorService�
Search*.aggregator.methods.SearchCreationsRequest+.aggregator.methods.SearchCreationsResponse")���#!/api/search/videos/{title}/{page}w
Login .aggregator.methods.LoginRequest!.aggregator.methods.LoginResponse")���#"/api/user/login:user_credentials�
WatchCreation(.aggregator.methods.WatchCreationRequest).aggregator.methods.WatchCreationResponse" ���/api/watch/{creation_id}�
SimilarCreations+.aggregator.methods.SimilarCreationsRequest$.aggregator.methods.GetCardsResponse"(���" /api/watch/similar/{creation_id}�
InitialComments*.aggregator.methods.InitialCommentsRequest+.aggregator.methods.InitialCommentsResponse"1���+)/api/watch/comments/{request.creation_id}�
GetTopComments).aggregator.methods.GetTopCommentsRequest*.aggregator.methods.GetTopCommentsResponse"D���></api/watch/comments/{request.creation_id=*}/{request.page=*}�
GetSecondComments,.aggregator.methods.GetSecondCommentsRequest-.aggregator.methods.GetSecondCommentsResponse"\���VT/api/watch/comments/second/{request.creation_id=*}/{request.root=*}/{request.page=*}�
GetUserReviews%.aggregator.methods.GetReviewsRequest*.aggregator.methods.GetUserReviewsResponse"!���"/api/review/query/user:*�
GetCreationReviews%.aggregator.methods.GetReviewsRequest..aggregator.methods.GetCreationReviewsResponse"%���"/api/review/query/creation:*�
GetCommentReviews%.aggregator.methods.GetReviewsRequest-.aggregator.methods.GetCommentReviewsResponse"$���"/api/review/query/comment:*�
GetNewUserReviews(.aggregator.methods.GetNewReviewsRequest*.aggregator.methods.GetUserReviewsResponse"%���"/api/review/query/new/user:*�
GetNewCreationReviews(.aggregator.methods.GetNewReviewsRequest..aggregator.methods.GetCreationReviewsResponse")���#"/api/review/query/new/creation:*�
GetNewCommentReviews(.aggregator.methods.GetNewReviewsRequest-.aggregator.methods.GetCommentReviewsResponse"(���""/api/review/query/new/comment:*m
HomePage.aggregator.methods.HomeRequest$.aggregator.methods.GetCardsResponse"���"/api/home/fetch:*~
Collections&.aggregator.methods.CollectionsRequest$.aggregator.methods.GetCardsResponse"!���"/api/collections/fetch:*r
History".aggregator.methods.HistoryRequest$.aggregator.methods.GetCardsResponse"���"/api/history/fetch:*B;Z9github.com/Yux77Yux/platform_backend/generated/aggregatorJ�
  

  

 

 P
	
 P
	
  &
	
 &
	
	 )
	

 (
	
 (


  


 

  

  

  6

  Ak

  

	  �ʼ"

 	 User OK


 

 +

 6V

 

	 �ʼ"
!
 !  WatchCreation OK


 

 ;

 Fn

  

	 �ʼ" 

 $( 相似视频 OK


 $

 $A

 $Lo

 %'

	 �ʼ"%'

 +/ Comment OK


 +

 +?

 +Jt

 ,.

	 �ʼ",.

 15

 1

 1=

 1Hq

 24

	 �ʼ"24

 7;

 7

 7C

 7Nz

 8:

	 �ʼ"8:

 >C Review


 >

 >9

 >Dm

 ?B

	 �ʼ"?B

 EJ

 E

 E=

 EHu

 FI

	 �ʼ"FI

 	LQ

 	L

 	L<

 	LGs

 	MP

	 	�ʼ"MP

 
SX

 
S

 
S?

 
SJs

 
TW

	 
�ʼ"TW

 Z_

 Z

 ZC

 ZN{

 [^

	 �ʼ"[^

 af

 a

 aB

 aMy

 be

	 �ʼ"be

 in 主页


 i

 i-

 i8[

 jm

	 �ʼ"jm

 qv 收藏夹


 q

 q7

 qBe

 ru

	 �ʼ"ru

 y~ 历史


 y

 y/

 y:]

 z}

	 �ʼ"z}bproto3
�
user/messages/user_auth.protouser.messagesuser/messages/user_role.proto"Y
UserAuth
user_id (RuserId4
	user_role (2.user.messages.UserRoleRuserRoleB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J
	
  '


  


 

  	

  	

  	

  	

 
'

 


 
"

 
%&bproto3
�
auth/methods/login.protoauth.methodsauth/messages/tokens.protouser/messages/user_auth.protocommon/api_response.proto"D
LoginRequest4
	user_auth (2.user.messages.UserAuthRuserAuth"e
LoginResponse-
tokens (2.auth.messages.TokensRtokens%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
	
  $
	
 '
	
 #


 
 


 


  '

  

  "

  %&


 




 "

 

 

  !







bproto3
�
auth/methods/refresh.protoauth.methodscommon/access_token.proto!auth/messages/refresh_token.protocommon/api_response.proto"R
RefreshRequest@
refresh_token (2.auth.messages.RefreshTokenRrefreshToken"p
RefreshResponse6
access_token (2.common.AccessTokenRaccessToken%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
	
  #
	
 +
	
 #


 
 


 


  /

  

  *

  -.


 




 &

 

 !

 $%







bproto3
�
auth/methods/check.protoauth.methodscommon/access_token.protocommon/api_response.proto"F
CheckRequest6
access_token (2.common.AccessTokenRaccessToken"6
CheckResponse%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
	
  #
	
 #


 	 


 	

  
&

  


  
!

  
$%


 




 

 

 

 bproto3
�
auth/auth_service.protoauthgoogle/api/annotations.protoauth/methods/login.protoauth/methods/refresh.protoauth/methods/check.proto2�
AuthService@
Login.auth.methods.LoginRequest.auth.methods.LoginResponsep
Refresh.auth.methods.RefreshRequest.auth.methods.RefreshResponse"(���""/api/auth/refresh:refresh_token@
Check.auth.methods.CheckRequest.auth.methods.CheckResponseB5Z3github.com/Yux77Yux/platform_backend/generated/authJ�
  

  

 

 J
	
 J
	
  &
	
 "
	
	 $
	

 "


  


 

  N

  

  '

  2L

 

 

 +

 6R

 


	 �ʼ"


 N

 

 '

 2Lbproto3
�
google/protobuf/empty.protogoogle.protobuf"
EmptyB}
com.google.protobufB
EmptyProtoPZ.google.golang.org/protobuf/types/known/emptypb��GPB�Google.Protobuf.WellKnownTypesJ�
 2
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" E
	
" E

# ,
	
# ,

$ +
	
$ +

% "
	

% "

& !
	
$& !

' ;
	
%' ;

( 
	
( 
�
 2 � A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }




 2bproto3
�
user/messages/follow.protouser.messages"J
Follow
follower_id (R
followerId
followee_id (R
followeeIdB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  	

  

 

 J
	
 J


  	


 

  

  

  

  

 

 

 

 bproto3
�
user/methods/post.protouser.methods$user/messages/user_credentials.protouser/messages/follow.protocommon/api_response.protocommon/access_token.proto"v
FollowRequest-
follow (2.user.messages.FollowRfollow6
access_token (2.common.AccessTokenRaccessToken"7
FollowResponse%
msg (2.common.ApiResponseRmsg"\
RegisterRequestI
user_credentials (2.user.messages.UserCredentialsRuserCredentials"9
RegisterResponse%
msg (2.common.ApiResponseRmsg"�
AddReviewerRequestI
user_credentials (2.user.messages.UserCredentialsRuserCredentials6
access_token (2.common.AccessTokenRaccessToken"<
AddReviewerResponse%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  #

  

 

 J
	
 J
	
  .
	
 $
	
 #
	
	 #


  


 

  "

  

  

   !

 &

 

 !

 $%


 




 

 

 

 


 




 5

 

  0

 34


 




 

 

 

 


 




 5

 

  0

 34

&



!

$%


! #


!

 "

 "

 "

 "bproto3
�
user/messages/user_gender.protouser.messages*1

UserGender
	UNDEFINED 
MALE

FEMALEB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  


  

 

 J
	
 J


  



 

  

  

  

 

 

 	


 	

 	

 	bproto3
�
user/messages/user_status.protouser.messages*K

UserStatus

HIDING 
INACTIVE

ACTIVE
LIMITED

DELETEB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J


  


 

  

  

  

 

 


 

 	

 	

 	

 


 
	

 


 

 

 bproto3
�
user/messages/user.protouser.messagesgoogle/protobuf/timestamp.protouser/messages/user_role.protouser/messages/user_gender.protouser/messages/user_status.protocommon/user_default.protocommon/custom_options.proto"�
User6
user_default (2.common.UserDefaultRuserDefault
user_avatar (	R
userAvatar 
user_bio (	B��RuserBio:
user_status (2.user.messages.UserStatusR
userStatus:
user_gender (2.user.messages.UserGenderR
userGender4
	user_role (2.user.messages.UserRoleRuserRole7
	user_bday (2.google.protobuf.TimestampRuserBdayB
user_created_at (2.google.protobuf.TimestampRuserCreatedAtB
user_updated_at	 (2.google.protobuf.TimestampRuserUpdatedAt
	followers
 (R	followers
	followees (R	followees"9
AnyUser.
any_user (2.user.messages.UserRanyUserB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
   

  

 

 J
	
 J
	
  )
	
 '
	
	 )
	

 )
	
 #
	
 %


  


 

  &

  

  !

  $%

 

 

 	

 

 <

 

 	

 

 ;

 ܆:

 +

 

 &

 )*

 +

 

 &

 )*

 '

 

 "

 %&

 *

 

 %

 ()

 0

 

 +

 ./

 0

 

 +

 ./

 	" 粉丝数量


 	

 	

 	

 
" 关注数量


 


 


 



  




 

 


 

 

 bproto3
�
user/methods/get.protouser.methodsuser/messages/follow.protouser/messages/user.protouser/messages/user_login.proto$user/messages/user_credentials.proto"common/user_creation_comment.protocommon/api_response.proto"?
GetFollowRequest
user_id (RuserId
page (Rpage"�
GetFollowResponse%
msg (2.common.ApiResponseRmsg1
users (2.common.UserCreationCommentRusers
master (Rmaster"#
GetUsersRequest
ids (Rids"l
GetUsersResponse1
users (2.common.UserCreationCommentRusers%
msg (2.common.ApiResponseRmsg"E
ExistFolloweeRequest-
follow (2.user.messages.FollowRfollow"T
ExistFolloweeResponse
exist (Rexist%
msg (2.common.ApiResponseRmsg")
GetUserRequest
user_id (RuserId"a
GetUserResponse'
user (2.user.messages.UserRuser%
msg (2.common.ApiResponseRmsg"Y
LoginRequestI
user_credentials (2.user.messages.UserCredentialsRuserCredentials"o
LoginResponse7

user_login (2.user.messages.UserLoginR	userLogin%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�

  <

  

 

 J
	
 J
	
  $
	
 "
	
 (
	
	 .
	
 ,
	
 #


  


 

  

  

  

  

 

 

 

 


 




 

 

 

 

0




%

&+

./










 




 

 


 

 

 


 !




 0

 


 %

 &+

 ./

 

 

 

 


# %


#

 $"

 $

 $

 $ !


' *


'

 (

 (

 (

 (

)

)

)

)


, .


,

 -

 -

 -

 -


0 3


0

 1

 1

 1

 1

2

2

2

2


5 7


5

 67

 6!

 6"2

 656


	9 <


	9

	 :+

	 :

	 :&

	 :)*

	;

	;

	;

	;bproto3
�
%user/messages/user_update_space.protouser.messagesgoogle/protobuf/timestamp.protouser/messages/user_gender.protocommon/user_default.proto"�
UserUpdateSpace6
user_default (2.common.UserDefaultRuserDefault
user_bio (	RuserBio:
user_gender (2.user.messages.UserGenderR
userGender7
	user_bday (2.google.protobuf.TimestampRuserBdayB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J
	
  )
	
 )
	
	 #


  


 

  &

  

  !

  $%

 

 

 	

 

 +

 

 &

 )*

 *

 

 %

 ()bproto3
�
&user/messages/user_update_status.protouser.messagesuser/messages/user_status.proto"g
UserUpdateStatus
user_id (RuserId:
user_status (2.user.messages.UserStatusR
userStatusB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  

  

 

 J
	
 J
	
  )


  


 

  	

  	

  	

  	

 
+

 


 
&

 
)*bproto3
�
&user/messages/user_update_avatar.protouser.messages"L
UserUpdateAvatar
user_id (RuserId
user_avatar (	R
userAvatarB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  	

  

 

 J
	
 J


  	


 

  

  

  

  

 

 

 	

 bproto3
�
#user/messages/user_update_bio.protouser.messages"C
UserUpdateBio
user_id (RuserId
user_bio (	RuserBioB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  	

  

 

 J
	
 J


  	


 

  

  

  

  

 

 

 	

 bproto3
�
user/methods/update.protouser.methods%user/messages/user_update_space.proto&user/messages/user_update_status.proto&user/messages/user_update_avatar.proto#user/messages/user_update_bio.protocommon/access_token.protocommon/api_response.proto"m
DelReviewerRequest
reviewer_id (R
reviewerId6
access_token (2.common.AccessTokenRaccessToken"<
DelReviewerResponse%
msg (2.common.ApiResponseRmsg"�
UpdateUserSpaceRequestJ
user_update_space (2.user.messages.UserUpdateSpaceRuserUpdateSpace6
access_token (2.common.AccessTokenRaccessToken"�
UpdateUserAvatarRequestM
user_update_avatar (2.user.messages.UserUpdateAvatarRuserUpdateAvatar6
access_token (2.common.AccessTokenRaccessToken"�
UpdateUserAvatarResponse%
msg (2.common.ApiResponseRmsgM
user_update_avatar (2.user.messages.UserUpdateAvatarRuserUpdateAvatar"�
UpdateUserStatusRequestM
user_update_status (2.user.messages.UserUpdateStatusRuserUpdateStatus6
access_token (2.common.AccessTokenRaccessToken"�
UpdateUserBioRequestD
user_update_bio (2.user.messages.UserUpdateBioRuserUpdateBio6
access_token (2.common.AccessTokenRaccessToken";
UpdateUserResponse%
msg (2.common.ApiResponseRmsgB5Z3github.com/Yux77Yux/platform_backend/generated/userJ�
  9

  

 

 J
	
 J
	
  /
	
 0
	
 0
	
	 -
	

 #
	
 #


  


 

  

  

  

  

 &

 

 !

 $%


 




 

 

 

 


 




 6

 

  1

 45

&



!

$%
(
 ! avatar
2 单字段 请求





 8

  

 !3

 67

 &

 

 !

 $%

$ ' avatar



$ 

 %

 %

 %

 %

&8

& 

&!3

&67

+ . status



+

 ,8

 , 

 ,!3

 ,67

-&

-

-!

-$%

1 4 bio



1

 22

 2

 2-

 201

3&

3

3!

3$%

7 9 响应



7

 8

 8

 8

 8bproto3
�
user/user_service.protousergoogle/api/annotations.protogoogle/protobuf/empty.protouser/methods/post.protouser/methods/get.protouser/methods/update.proto2�
UserService{
AddReviewer .user.methods.AddReviewerRequest!.user.methods.AddReviewerResponse"'���!"/api/reviewer:user_credentials`
Follow.user.methods.FollowRequest.user.methods.FollowResponse"���"/api/user/follow:*w
Register.user.methods.RegisterRequest.user.methods.RegisterResponse",���&"/api/user/register:user_credentials`
CancelFollow.user.methods.FollowRequest.google.protobuf.Empty"���2/api/user/follow:*�
ExistFollowee".user.methods.ExistFolloweeRequest#.user.methods.ExistFolloweeResponse"F���@>/api/user/follow/get/{follow.followee_id}/{follow.follower_id}@
Login.user.methods.LoginRequest.user.methods.LoginResponse�
GetFolloweesByTime.user.methods.GetFollowRequest.user.methods.GetFollowResponse"0���*(/api/user/followee/time/{user_id}/{page}�
GetFolloweesByViews.user.methods.GetFollowRequest.user.methods.GetFollowResponse"1���+)/api/user/followee/views/{user_id}/{page}|
GetFollowers.user.methods.GetFollowRequest.user.methods.GetFollowResponse"+���%#/api/user/follower/{user_id}/{page}c
GetUser.user.methods.GetUserRequest.user.methods.GetUserResponse"���/api/user/{user_id}I
GetUsers.user.methods.GetUsersRequest.user.methods.GetUsersResponseu
UpdateUserSpace$.user.methods.UpdateUserSpaceRequest .user.methods.UpdateUserResponse"���2/api/user/space:*~
UpdateUserAvatar%.user.methods.UpdateUserAvatarRequest&.user.methods.UpdateUserAvatarResponse"���2/api/user/avatar:*x
UpdateUserStatus%.user.methods.UpdateUserStatusRequest .user.methods.UpdateUserResponse"���2/api/user/status:*o
UpdateUserBio".user.methods.UpdateUserBioRequest .user.methods.UpdateUserResponse"���2/api/user/bio:*q
DelReviewer .user.methods.DelReviewerRequest!.user.methods.DelReviewerResponse"���2/api/reviewer/auth:*B6Z4github.com/Yux77Yux/platform_backend/generated/user;J�
  u

  

 

 K
	
 K
	
  &
	
 %
	
	 !
	

  
	
 #


  u


 

   POST


  

  1

  <\

  

	  �ʼ"

 

 

 '

 2M

 

	 �ʼ"

 "

 

 +

 6S

 !

	 �ʼ"!

 %* DEL


 %

 %-

 %8M

 &)

	 �ʼ"&)

 -1 GET


 -

 -5

 -@b

 .0

	 �ʼ".0

 3L

 3

 3%

 30J

 59

 5

 56

 5A_

 68

	 �ʼ"68

 ;?

 ;

 ;7

 ;B`

 <>

	 �ʼ"<>

 AE

 A

 A0

 A;Y

 BD

	 �ʼ"BD

 	GK

 	G

 	G)

 	G4P

 	HJ

	 	�ʼ"HJ

 
LU

 
L

 
L+

 
L6S
$
 PU UPDATE
 批量字段


 P

 P9

 PDc

 QT

	 �ʼ"QT

 X] User Avatar


 X

 X;

 XFk

 Y\

	 �ʼ"Y\

 `e User Status


 `

 `;

 `Fe

 ad

	 �ʼ"ad

 hm
 User Bio


 h

 h5

 h@_

 il

	 �ʼ"il

 ot

 o

 o1

 o<\

 ps

	 �ʼ"psbproto3
�
creation/methods/post.protocreation.methodscommon/api_response.proto'creation/messages/creation_upload.protocommon/access_token.proto"�
UploadCreationRequest>
	base_info (2!.creation.messages.CreationUploadRbaseInfo6
access_token (2.common.AccessTokenRaccessToken"?
UploadCreationResponse%
msg (2.common.ApiResponseRmsgB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N
	
  #
	
 1
	
 #


 
 


 


  1

  "

  #,

  /0

 &

 

 !

 $%


 




 

 

 

 bproto3
�

'creation/messages/creation_update.protocreation.messages'creation/messages/creation_status.proto"�
CreationUpdated
creation_id (R
creationId
	thumbnail (	R	thumbnail
title (	Rtitle
bio (	Rbio
	author_id (RauthorId
src (	Rsrc
duration (Rduration9
status (2!.creation.messages.CreationStatusRstatus"�
CreationUpdateStatus
creation_id (R
creationId
	author_id (RauthorId9
status (2!.creation.messages.CreationStatusRstatusB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N
	
  1


  


 

  	

  	

  	

  	

 


 


 
	

 


 

 

 	

 

 

 

 	

 

 

 

 

 

 

 

 	

 

 

 

 

 

 .

 "

 #)

 ,-


 




 

 

 

 









.

"

#)

,-bproto3
�
creation/methods/update.protocreation.methodscommon/api_response.proto'creation/messages/creation_update.protocommon/access_token.proto"�
UpdateCreationRequestC
update_info (2".creation.messages.CreationUpdatedR
updateInfo6
access_token (2.common.AccessTokenRaccessToken"�
UpdateCreationStatusRequestH
update_info (2'.creation.messages.CreationUpdateStatusR
updateInfo6
access_token (2.common.AccessTokenRaccessToken"?
UpdateCreationResponse%
msg (2.common.ApiResponseRmsgB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N
	
  #
	
 1
	
 #


 
 


 


  4

  #

  $/

  23

 &

 

 !

 $%


 


#

 9

 (

 )4

 78

&



!

$%


 




 

 

 

 bproto3
�
creation/methods/get.protocreation.methodscommon/api_response.protocommon/access_token.proto creation/messages/creation.proto'creation/messages/creation_status.proto"5
GetCreationRequest
creation_id (R
creationId"t
GetCreationPrivateRequest
creation_id (R
creationId6
access_token (2.common.AccessTokenRaccessToken"�
GetCreationResponseD
creation_info (2.creation.messages.CreationInfoRcreationInfo%
msg (2.common.ApiResponseRmsg"�
GetUserCreationsRequest9
status (2!.creation.messages.CreationStatusRstatus
user_id (RuserId
page (Rpage6
access_token (2.common.AccessTokenRaccessToken"{
GetSpaceCreationsRequest
user_id (RuserId
page (Rpage2
by_what (2.creation.methods.ByCountRbyWhat"*
GetCreationListRequest
ids (Rids"�
GetCreationListResponse%
msg (2.common.ApiResponseRmsgO
creation_info_group (2.creation.messages.CreationInfoRcreationInfoGroup
count (Rcount"A
SearchCreationRequest
title (	Rtitle
page (Rpage*D
ByCount
PUBLISHED_TIME 	
VIEWS	
LIKES
COLLECTIONSB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  >

  

 

 N
	
 N
	
  #
	
 #
	
 *
	
	 1

   视频详细页



 

  

  

  

  


 


!

 

 

 

 

&



!

$%


 




 3

  

 !.

 12










  


 

  

  

  

 

 

 


 

 

 


 

 

 

" ' 视频管理



"

 #.

 #"

 ##)

 #,-

$

$

$

$

%

%

%

%

&&

&

&!

&$%

* . 空间



* 

 +

 +

 +

 +

,

,

,

,

-

-	

-


-
+
1 3 发布状态的Creation列表



1

 2

 2


 2

 2

 2


5 9


5

 6

 6

 6

 6

7B

7


7)

7*=

7@A

8

8

8

8


; >


;

 <

 <

 <	

 <

=

=

=

=bproto3
�
creation/methods/delete.protocreation.methodscommon/access_token.proto"p
DeleteCreationRequest
creation_id (R
creationId6
access_token (2.common.AccessTokenRaccessTokenB9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  

  

 

 N
	
 N
	
  #


  


 

  	

  	

  	

  	

 
&

 


 
!

 
$%bproto3
�
creation/creation_service.protocreationgoogle/api/annotations.protogoogle/protobuf/empty.protocreation/methods/post.protocreation/methods/update.protocreation/methods/get.protocreation/methods/delete.proto2�

CreationService}
UploadCreation'.creation.methods.UploadCreationRequest(.creation.methods.UploadCreationResponse"���"/api/creation:*Z
GetCreation$.creation.methods.GetCreationRequest%.creation.methods.GetCreationResponse�
GetCreationPrivate+.creation.methods.GetCreationPrivateRequest%.creation.methods.GetCreationResponse" ���"/api/creation/private:*�
GetSpaceCreations*.creation.methods.GetSpaceCreationsRequest).creation.methods.GetCreationListResponse"6���0./api/creation/space/{user_id}/{page}/{by_what}�
GetUserCreations).creation.methods.GetUserCreationsRequest).creation.methods.GetCreationListResponse"���"/api/manager/creationsd
SearchCreation'.creation.methods.SearchCreationRequest).creation.methods.GetCreationListResponsef
GetCreationList(.creation.methods.GetCreationListRequest).creation.methods.GetCreationListResponsel
GetPublicCreationList(.creation.methods.GetCreationListRequest).creation.methods.GetCreationListResponser
DeleteCreation'.creation.methods.DeleteCreationRequest.google.protobuf.Empty"���"/api/creation/delete:*}
UpdateCreation'.creation.methods.UpdateCreationRequest(.creation.methods.UpdateCreationResponse"���2/api/creation:*�
PublishDraftCreation-.creation.methods.UpdateCreationStatusRequest(.creation.methods.UpdateCreationResponse"���2/api/creation/status:*B9Z7github.com/Yux77Yux/platform_backend/generated/creationJ�
  F

  

 

 N
	
 N
	
  &
	
 %
	
	 %
	

 '
	
 $
	
 '


  F


 

   POST


  

  ;

  Fm

  

	  �ʼ"

 f GET


 

 5

 @d

 

 

 C

 Nr

 

	 �ʼ"

  $

  

  A

  Lt

 !#

	 �ʼ"!#

 &*

 &

 &?

 &Jr

 ')

	 �ʼ"')

 ,p

 ,

 ,;

 ,Fn

 -r

 -

 -=

 -Hp

 .x

 .

 .C

 .Nv

 16 DELETE


 1

 1;

 1F[

 25

	 �ʼ"25

 	9> UPDATE


 	9

 	9;

 	9Fm

 	:=

	 	�ʼ":=

 
@E

 
@

 
@G

 
@Ry

 
AD

	 
�ʼ"ADbproto3
�
comment/methods/post.protocomment.methodscommon/api_response.protocomment/messages/comment.protocommon/access_token.proto"�
PublishCommentRequest3
comment (2.comment.messages.CommentRcomment6
access_token (2.common.AccessTokenRaccessToken"?
PublishCommentResponse%
msg (2.common.ApiResponseRmsgB8Z6github.com/Yux77Yux/platform_backend/generated/commentJ�
  

  

 

 M
	
 M
	
  #
	
 (
	
 #


 
 


 


  '

  

  "

  %&

 &

 

 !

 $%


 




 

 

 

 bproto3
�
comment/methods/delete.protocomment.methodscommon/access_token.proto"�
DeleteCommentRequest

comment_id (R	commentId
creation_id (R
creationId6
access_token (2.common.AccessTokenRaccessTokenB8Z6github.com/Yux77Yux/platform_backend/generated/commentJ�
  

  

 

 M
	
 M
	
  #


  


 

  	

  	

  	

  	

 


 


 


 


 &

 

 !

 $%bproto3
�
comment/comment_service.protocommentgoogle/api/annotations.protogoogle/protobuf/empty.protocomment/methods/post.protocomment/methods/get.protocomment/methods/delete.proto2�
CommentServicez
PublishComment&.comment.methods.PublishCommentRequest'.comment.methods.PublishCommentResponse"���"/api/comment:*d
InitialComments'.comment.methods.InitialCommentsRequest(.comment.methods.InitialCommentsResponsen
DeleteComment%.comment.methods.DeleteCommentRequest.google.protobuf.Empty"���2/api/comment/delete:*a
GetTopComments&.comment.methods.GetTopCommentsRequest'.comment.methods.GetTopCommentsResponsej
GetSecondComments).comment.methods.GetSecondCommentsRequest*.comment.methods.GetSecondCommentsResponseb
GetReplyComments(.comment.methods.GetReplyCommentsRequest$.comment.methods.GetCommentsResponseU

GetComment".comment.methods.GetCommentRequest#.comment.methods.GetCommentResponseX
GetComments#.comment.methods.GetCommentsRequest$.comment.methods.GetCommentsResponseB8Z6github.com/Yux77Yux/platform_backend/generated/commentJ�
  '

  

 

 M
	
 M
	
  &
	
 %
	
	 $
	

 #
	
 &


  '


 

  

  

  :

  Ek

  

	  �ʼ"

 p

 

 <

 Gn

 

 

 8

 CX

 

	 �ʼ"

 m

 

 :

 Ek

  v

  

  @

  Kt

 "n

 "

 ">

 "Il

 $a

 $

 $2

 $=_

 &d

 &

 &4

 &?bbproto3
�
review/methods/update.protoreview.methodsreview/messages/review.protocommon/access_token.protocommon/api_response.proto"m
DelReviewerRequest
reviewer_id (R
reviewerId6
access_token (2.common.AccessTokenRaccessToken"<
DelReviewerResponse%
msg (2.common.ApiResponseRmsg"~
UpdateReviewRequest/
review (2.review.messages.ReviewRreview6
access_token (2.common.AccessTokenRaccessToken"=
UpdateReviewResponse%
msg (2.common.ApiResponseRmsgB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L
	
  &
	
 #
	
 #


 
 


 


  

  

  

  

 &

 

 !

 $%


 




 

 

 

 


 




 $

 

 

 "#

&



!

$%


 




 

 

 

 bproto3
�
review/methods/get.protoreview.methodsreview/messages/status.protoreview/messages/type.protoreview/messages/review.protocommon/api_response.proto"�
GetReviewsRequest/
type (2.review.messages.TargetTypeRtype5
status (2.review.messages.ReviewStatusRstatus
page (Rpage
reviewer_id (R
reviewerId"h
GetNewReviewsRequest
reviewer_id (R
reviewerId/
type (2.review.messages.TargetTypeRtype"�
GetReviewsResponse%
msg (2.common.ApiResponseRmsg1
reviews (2.review.messages.ReviewRreviews
count (Rcount"(
GetReviewDetailRequest
id (Rid"q
GetReviewDetailResponse%
msg (2.common.ApiResponseRmsg/
review (2.review.messages.ReviewRreviewB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  %

  

 

 L
	
 L
	
  &
	
 $
	
 &
	
	 #


  


 
6
  &") 获取类型，评论，用户，作品


  

  !

  $%
9
 *", 获取需要审核的，或已被审核的


 

 %

 ()

 

 

 

 

 

 

 

 


 




 

 

 

 

&



!

$%


 




 

 

 

 

.




!

")

,-








8
  , 用于给 用户 查看 是否审核完毕





 

 

 


 


" %


"

 #

 #

 #

 #

$$

$

$

$"#bproto3
�
review/methods/post.protoreview.methods review/messages/new_review.protocommon/api_response.proto"@
NewReviewRequest,
new (2.review.messages.NewReviewRnew":
NewReviewResponse%
msg (2.common.ApiResponseRmsgB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L
	
  *
	
 #


 	 


 	

  
&

  


  
!

  
$%


 




 

 

 

 bproto3
�	
review/review_service.protoreviewgoogle/api/annotations.protoreview/methods/update.protoreview/methods/get.protoreview/methods/post.proto2�
ReviewServiceu
UpdateReview#.review.methods.UpdateReviewRequest$.review.methods.UpdateReviewResponse"���2/api/review/updatee
	NewReview .review.methods.NewReviewRequest!.review.methods.NewReviewResponse"���"/api/review\
	GetReview&.review.methods.GetReviewDetailRequest'.review.methods.GetReviewDetailResponseS

GetReviews!.review.methods.GetReviewsRequest".review.methods.GetReviewsResponseY
GetNewReviews$.review.methods.GetNewReviewsRequest".review.methods.GetReviewsResponseB7Z5github.com/Yux77Yux/platform_backend/generated/reviewJ�
  

  

 

 L
	
 L
	
  &
	
 %
	
	 "
	

 #


  


 

   UPDATE


  

  5

  @c

  

	  �ʼ"

  POST


 

 /

 :Z

 

	 �ʼ"

 h GET


 

 5

 @f

 _

 

 1

 <]

 e

 

 7

 Bcbproto3
�
+interaction/messages/base_interaction.protointeraction.messagescommon/operate.protogoogle/protobuf/timestamp.proto"K
BaseInteraction
user_id (RuserId
creation_id (R
creationId"�
OperateInteraction'
action (2.common.OperateRaction9
base (2%.interaction.messages.BaseInteractionRbase9

updated_at (2.google.protobuf.TimestampR	updatedAt3
save_at (2.google.protobuf.TimestampRsaveAt"}
OperateAnyInteraction'
action (2.common.OperateRaction;
bases (2%.interaction.messages.BaseInteractionRbases"t
AnyOperateInteraction[
operate_interactions (2(.interaction.messages.OperateInteractionRoperateInteractionsB<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�
  

  

 

 Q
	
 Q
	
  
	
 )


 	 


 	

  


  


  


  


 

 

 

 


 




 

 

 

 









+



&

)*

(



#

&'


 




 

 

 

 

%






 

#$


 




 7

 


 

 2

 56bproto3
�
interaction/methods/post.protointeraction.methods+interaction/messages/base_interaction.protocommon/access_token.protocommon/api_response.proto"�
PostInteractionRequest9
base (2%.interaction.messages.BaseInteractionRbase6
access_token (2.common.AccessTokenRaccessToken"@
PostInteractionResponse%
msg (2.common.ApiResponseRmsgB<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�
  

  

 

 Q
	
 Q
	
  5
	
 #
	
	 #


  


 

  0

  &

  '+

  ./

 &

 

 !

 $%


 




 

 

 

 bproto3
�
&interaction/messages/interaction.protointeraction.messages+interaction/messages/base_interaction.protogoogle/protobuf/timestamp.proto"�
Interaction9
base (2%.interaction.messages.BaseInteractionRbase

action_tag (R	actionTag9

updated_at (2.google.protobuf.TimestampR	updatedAt3
save_at (2.google.protobuf.TimestampRsaveAt"Z
AnyInteractionH
any_interction (2!.interaction.messages.InteractionRanyInterctionB<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�
  

  

 

 Q
	
 Q
	
  5
	
 )


 
 


 


  0

  &

  '+

  ./

 

 

 

 

 +

 

 &

 )*

 (

 

 #

 &'


 




 *

 


 

 %

 ()bproto3
�
interaction/methods/get.protointeraction.methods&interaction/messages/interaction.proto+interaction/messages/base_interaction.protocommon/access_token.protocommon/api_response.proto"%
GetRecommendRequest
id (Rid"[
GetRecommendResponse%
msg (2.common.ApiResponseRmsg
	creations (R	creations"�
GetCreationInteractionRequest9
base (2%.interaction.messages.BaseInteractionRbase6
access_token (2.common.AccessTokenRaccessToken"f
GetCreationInteractionResponse%
msg (2.common.ApiResponseRmsg

action_tag (R	actionTag"B
GetHistoriesRequest
user_id (RuserId
page (Rpage"D
GetCollectionsRequest
user_id (RuserId
page (Rpage"�
GetInteractionsResponse%
msg (2.common.ApiResponseRmsgM
any_interaction (2$.interaction.messages.AnyInteractionRanyInteractionB<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�
  ,

  

 

 Q
	
 Q
	
  0
	
 5
	
	 #
	

 #


  


 

  

  

  


  


 




 

 

 

 













 


%

 0

 &

 '+

 ./

&



!

$%


 


&

 

 

 

 










 "




  

  

  

  

!

!

!

!


$ '


$

 %

 %

 %

 %

&

&

&

&


) ,


)

 *

 *

 *

 *

+:

+%

+&5

+89bproto3
�
 interaction/methods/update.protointeraction.methods+interaction/messages/base_interaction.protocommon/api_response.protocommon/access_token.proto"�
UpdateInteractionRequest9
base (2%.interaction.messages.BaseInteractionRbase6
access_token (2.common.AccessTokenRaccessToken"�
UpdateInteractionsRequest;
bases (2%.interaction.messages.BaseInteractionRbases6
access_token (2.common.AccessTokenRaccessToken"B
UpdateInteractionResponse%
msg (2.common.ApiResponseRmsgB<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�
  

  

 

 Q
	
 Q
	
  5
	
 #
	
	 #


  


  

  0

  &

  '+

  ./

 &

 

 !

 $%


 


!

 :

 


 /

 05

 89

&



!

$%


 


!

 

 

 

 bproto3
�
%interaction/interaction_service.protointeractiongoogle/api/annotations.protointeraction/methods/post.protointeraction/methods/get.proto interaction/methods/update.proto2�
InteractionServicel
PostInteraction+.interaction.methods.PostInteractionRequest,.interaction.methods.PostInteractionResponsek
GetRecommendBaseUser(.interaction.methods.GetRecommendRequest).interaction.methods.GetRecommendResponseo
GetRecommendBaseCreation(.interaction.methods.GetRecommendRequest).interaction.methods.GetRecommendResponse�
GetActionTag2.interaction.methods.GetCreationInteractionRequest3.interaction.methods.GetCreationInteractionResponse"���2/api/interaction:*f
GetHistories(.interaction.methods.GetHistoriesRequest,.interaction.methods.GetInteractionsResponsej
GetCollections*.interaction.methods.GetCollectionsRequest,.interaction.methods.GetInteractionsResponse�
ClickCollection-.interaction.methods.UpdateInteractionRequest..interaction.methods.UpdateInteractionResponse"&��� 2/api/interaction/collection:*�
	ClickLike-.interaction.methods.UpdateInteractionRequest..interaction.methods.UpdateInteractionResponse" ���2/api/interaction/like:*�
CancelCollections..interaction.methods.UpdateInteractionsRequest..interaction.methods.UpdateInteractionResponse"-���'2"/api/interaction/collection/cancel:*�
DelHistories..interaction.methods.UpdateInteractionsRequest..interaction.methods.UpdateInteractionResponse"*���$2/api/interaction/history/cancel:*�

CancelLike-.interaction.methods.UpdateInteractionRequest..interaction.methods.UpdateInteractionResponse"'���!2/api/interaction/like/cancel:*B<Z:github.com/Yux77Yux/platform_backend/generated/interactionJ�	
  F

  

 

 Q
	
 Q
	
  &
	
 (
	
	 '
	

 *


  F


 
i
  x\ POST
 这个为展示完视频信息之后再查看是否登录，判断是否发送事件


  

  @

  Kv

 w Get


 

 B

 Mu

 {

 

 F

 Qy
j
 \ 这个为展示完视频信息之后再查看是否登录，再盘点是否已经点赞过


 

 D

 O�

 

	 �ʼ"

 r

 

 :

 Ep

 v

 

 >

 It
:
 $)2, UPDATE
 实际是修改对于的action tag


 $

 $B

 $Mz

 %(

	 �ʼ"%(

 +0

 +

 +<

 +Gt

 ,/

	 �ʼ",/

 27

 2

 2E

 2P}

 36

	 �ʼ"36

 	9>

 	9

 	9@

 	9Kx

 	:=

	 	�ʼ":=

 
@E

 
@

 
@=

 
@Hu

 
AD

	 
�ʼ"ADbproto3