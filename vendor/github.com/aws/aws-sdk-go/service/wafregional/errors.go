// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

const (

	// ErrCodeWAFDisallowedNameException for service response error code
	// "WAFDisallowedNameException".
	//
	// The name specified is invalid.
	ErrCodeWAFDisallowedNameException = "WAFDisallowedNameException"

	// ErrCodeWAFInternalErrorException for service response error code
	// "WAFInternalErrorException".
	//
	// The operation failed because of a system problem, even though the request
	// was valid. Retry your request.
	ErrCodeWAFInternalErrorException = "WAFInternalErrorException"

	// ErrCodeWAFInvalidAccountException for service response error code
	// "WAFInvalidAccountException".
	//
	// The operation failed because you tried to create, update, or delete an object
	// by using an invalid account identifier.
	ErrCodeWAFInvalidAccountException = "WAFInvalidAccountException"

	// ErrCodeWAFInvalidOperationException for service response error code
	// "WAFInvalidOperationException".
	//
	// The operation failed because there was nothing to do. For example:
	//
	//    * You tried to remove a Rule from a WebACL, but the Rule isn't in the
	//    specified WebACL.
	//
	//    * You tried to remove an IP address from an IPSet, but the IP address
	//    isn't in the specified IPSet.
	//
	//    * You tried to remove a ByteMatchTuple from a ByteMatchSet, but the ByteMatchTuple
	//    isn't in the specified WebACL.
	//
	//    * You tried to add a Rule to a WebACL, but the Rule already exists in
	//    the specified WebACL.
	//
	//    * You tried to add an IP address to an IPSet, but the IP address already
	//    exists in the specified IPSet.
	//
	//    * You tried to add a ByteMatchTuple to a ByteMatchSet, but the ByteMatchTuple
	//    already exists in the specified WebACL.
	ErrCodeWAFInvalidOperationException = "WAFInvalidOperationException"

	// ErrCodeWAFInvalidParameterException for service response error code
	// "WAFInvalidParameterException".
	//
	// The operation failed because AWS WAF didn't recognize a parameter in the
	// request. For example:
	//
	//    * You specified an invalid parameter name.
	//
	//    * You specified an invalid value.
	//
	//    * You tried to update an object (ByteMatchSet, IPSet, Rule, or WebACL)
	//    using an action other than INSERT or DELETE.
	//
	//    * You tried to create a WebACL with a DefaultActionType other than ALLOW,
	//    BLOCK, or COUNT.
	//
	//    * You tried to create a RateBasedRule with a RateKey value other than
	//    IP.
	//
	//    * You tried to update a WebACL with a WafActionType other than ALLOW,
	//    BLOCK, or COUNT.
	//
	//    * You tried to update a ByteMatchSet with a FieldToMatchType other than
	//    HEADER, METHOD, QUERY_STRING, URI, or BODY.
	//
	//    * You tried to update a ByteMatchSet with a Field of HEADER but no value
	//    for Data.
	//
	//    * Your request references an ARN that is malformed, or corresponds to
	//    a resource with which a web ACL cannot be associated.
	ErrCodeWAFInvalidParameterException = "WAFInvalidParameterException"

	// ErrCodeWAFInvalidRegexPatternException for service response error code
	// "WAFInvalidRegexPatternException".
	//
	// The regular expression (regex) you specified in RegexPatternString is invalid.
	ErrCodeWAFInvalidRegexPatternException = "WAFInvalidRegexPatternException"

	// ErrCodeWAFLimitsExceededException for service response error code
	// "WAFLimitsExceededException".
	//
	// The operation exceeds a resource limit, for example, the maximum number of
	// WebACL objects that you can create for an AWS account. For more information,
	// see Limits (http://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
	// in the AWS WAF Developer Guide.
	ErrCodeWAFLimitsExceededException = "WAFLimitsExceededException"

	// ErrCodeWAFNonEmptyEntityException for service response error code
	// "WAFNonEmptyEntityException".
	//
	// The operation failed because you tried to delete an object that isn't empty.
	// For example:
	//
	//    * You tried to delete a WebACL that still contains one or more Rule objects.
	//
	//    * You tried to delete a Rule that still contains one or more ByteMatchSet
	//    objects or other predicates.
	//
	//    * You tried to delete a ByteMatchSet that contains one or more ByteMatchTuple
	//    objects.
	//
	//    * You tried to delete an IPSet that references one or more IP addresses.
	ErrCodeWAFNonEmptyEntityException = "WAFNonEmptyEntityException"

	// ErrCodeWAFNonexistentContainerException for service response error code
	// "WAFNonexistentContainerException".
	//
	// The operation failed because you tried to add an object to or delete an object
	// from another object that doesn't exist. For example:
	//
	//    * You tried to add a Rule to or delete a Rule from a WebACL that doesn't
	//    exist.
	//
	//    * You tried to add a ByteMatchSet to or delete a ByteMatchSet from a Rule
	//    that doesn't exist.
	//
	//    * You tried to add an IP address to or delete an IP address from an IPSet
	//    that doesn't exist.
	//
	//    * You tried to add a ByteMatchTuple to or delete a ByteMatchTuple from
	//    a ByteMatchSet that doesn't exist.
	ErrCodeWAFNonexistentContainerException = "WAFNonexistentContainerException"

	// ErrCodeWAFNonexistentItemException for service response error code
	// "WAFNonexistentItemException".
	//
	// The operation failed because the referenced object doesn't exist.
	ErrCodeWAFNonexistentItemException = "WAFNonexistentItemException"

	// ErrCodeWAFReferencedItemException for service response error code
	// "WAFReferencedItemException".
	//
	// The operation failed because you tried to delete an object that is still
	// in use. For example:
	//
	//    * You tried to delete a ByteMatchSet that is still referenced by a Rule.
	//
	//    * You tried to delete a Rule that is still referenced by a WebACL.
	ErrCodeWAFReferencedItemException = "WAFReferencedItemException"

	// ErrCodeWAFStaleDataException for service response error code
	// "WAFStaleDataException".
	//
	// The operation failed because you tried to create, update, or delete an object
	// by using a change token that has already been used.
	ErrCodeWAFStaleDataException = "WAFStaleDataException"

	// ErrCodeWAFSubscriptionNotFoundException for service response error code
	// "WAFSubscriptionNotFoundException".
	//
	// The specified subscription does not exist.
	ErrCodeWAFSubscriptionNotFoundException = "WAFSubscriptionNotFoundException"

	// ErrCodeWAFUnavailableEntityException for service response error code
	// "WAFUnavailableEntityException".
	//
	// The operation failed because the entity referenced is temporarily unavailable.
	// Retry your request.
	ErrCodeWAFUnavailableEntityException = "WAFUnavailableEntityException"
)
