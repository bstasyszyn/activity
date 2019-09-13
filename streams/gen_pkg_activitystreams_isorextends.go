package streams

import (
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typepublickey "github.com/go-fed/activity/streams/impl/activitystreams/type_publickey"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// IsOrExtendsActivityStreamsAccept returns true if the other provided type is the
// Accept type or extends from the Accept type.
func IsOrExtendsActivityStreamsAccept(other vocab.Type) bool {
	return typeaccept.IsOrExtendsAccept(other)
}

// IsOrExtendsActivityStreamsActivity returns true if the other provided type is
// the Activity type or extends from the Activity type.
func IsOrExtendsActivityStreamsActivity(other vocab.Type) bool {
	return typeactivity.IsOrExtendsActivity(other)
}

// IsOrExtendsActivityStreamsAdd returns true if the other provided type is the
// Add type or extends from the Add type.
func IsOrExtendsActivityStreamsAdd(other vocab.Type) bool {
	return typeadd.IsOrExtendsAdd(other)
}

// IsOrExtendsActivityStreamsAnnounce returns true if the other provided type is
// the Announce type or extends from the Announce type.
func IsOrExtendsActivityStreamsAnnounce(other vocab.Type) bool {
	return typeannounce.IsOrExtendsAnnounce(other)
}

// IsOrExtendsActivityStreamsApplication returns true if the other provided type
// is the Application type or extends from the Application type.
func IsOrExtendsActivityStreamsApplication(other vocab.Type) bool {
	return typeapplication.IsOrExtendsApplication(other)
}

// IsOrExtendsActivityStreamsArrive returns true if the other provided type is the
// Arrive type or extends from the Arrive type.
func IsOrExtendsActivityStreamsArrive(other vocab.Type) bool {
	return typearrive.IsOrExtendsArrive(other)
}

// IsOrExtendsActivityStreamsArticle returns true if the other provided type is
// the Article type or extends from the Article type.
func IsOrExtendsActivityStreamsArticle(other vocab.Type) bool {
	return typearticle.IsOrExtendsArticle(other)
}

// IsOrExtendsActivityStreamsAudio returns true if the other provided type is the
// Audio type or extends from the Audio type.
func IsOrExtendsActivityStreamsAudio(other vocab.Type) bool {
	return typeaudio.IsOrExtendsAudio(other)
}

// IsOrExtendsActivityStreamsBlock returns true if the other provided type is the
// Block type or extends from the Block type.
func IsOrExtendsActivityStreamsBlock(other vocab.Type) bool {
	return typeblock.IsOrExtendsBlock(other)
}

// IsOrExtendsActivityStreamsCollection returns true if the other provided type is
// the Collection type or extends from the Collection type.
func IsOrExtendsActivityStreamsCollection(other vocab.Type) bool {
	return typecollection.IsOrExtendsCollection(other)
}

// IsOrExtendsActivityStreamsCollectionPage returns true if the other provided
// type is the CollectionPage type or extends from the CollectionPage type.
func IsOrExtendsActivityStreamsCollectionPage(other vocab.Type) bool {
	return typecollectionpage.IsOrExtendsCollectionPage(other)
}

// IsOrExtendsActivityStreamsCreate returns true if the other provided type is the
// Create type or extends from the Create type.
func IsOrExtendsActivityStreamsCreate(other vocab.Type) bool {
	return typecreate.IsOrExtendsCreate(other)
}

// IsOrExtendsActivityStreamsDelete returns true if the other provided type is the
// Delete type or extends from the Delete type.
func IsOrExtendsActivityStreamsDelete(other vocab.Type) bool {
	return typedelete.IsOrExtendsDelete(other)
}

// IsOrExtendsActivityStreamsDislike returns true if the other provided type is
// the Dislike type or extends from the Dislike type.
func IsOrExtendsActivityStreamsDislike(other vocab.Type) bool {
	return typedislike.IsOrExtendsDislike(other)
}

// IsOrExtendsActivityStreamsDocument returns true if the other provided type is
// the Document type or extends from the Document type.
func IsOrExtendsActivityStreamsDocument(other vocab.Type) bool {
	return typedocument.IsOrExtendsDocument(other)
}

// IsOrExtendsActivityStreamsEvent returns true if the other provided type is the
// Event type or extends from the Event type.
func IsOrExtendsActivityStreamsEvent(other vocab.Type) bool {
	return typeevent.IsOrExtendsEvent(other)
}

// IsOrExtendsActivityStreamsFlag returns true if the other provided type is the
// Flag type or extends from the Flag type.
func IsOrExtendsActivityStreamsFlag(other vocab.Type) bool {
	return typeflag.IsOrExtendsFlag(other)
}

// IsOrExtendsActivityStreamsFollow returns true if the other provided type is the
// Follow type or extends from the Follow type.
func IsOrExtendsActivityStreamsFollow(other vocab.Type) bool {
	return typefollow.IsOrExtendsFollow(other)
}

// IsOrExtendsActivityStreamsGroup returns true if the other provided type is the
// Group type or extends from the Group type.
func IsOrExtendsActivityStreamsGroup(other vocab.Type) bool {
	return typegroup.IsOrExtendsGroup(other)
}

// IsOrExtendsActivityStreamsIgnore returns true if the other provided type is the
// Ignore type or extends from the Ignore type.
func IsOrExtendsActivityStreamsIgnore(other vocab.Type) bool {
	return typeignore.IsOrExtendsIgnore(other)
}

// IsOrExtendsActivityStreamsImage returns true if the other provided type is the
// Image type or extends from the Image type.
func IsOrExtendsActivityStreamsImage(other vocab.Type) bool {
	return typeimage.IsOrExtendsImage(other)
}

// IsOrExtendsActivityStreamsIntransitiveActivity returns true if the other
// provided type is the IntransitiveActivity type or extends from the
// IntransitiveActivity type.
func IsOrExtendsActivityStreamsIntransitiveActivity(other vocab.Type) bool {
	return typeintransitiveactivity.IsOrExtendsIntransitiveActivity(other)
}

// IsOrExtendsActivityStreamsInvite returns true if the other provided type is the
// Invite type or extends from the Invite type.
func IsOrExtendsActivityStreamsInvite(other vocab.Type) bool {
	return typeinvite.IsOrExtendsInvite(other)
}

// IsOrExtendsActivityStreamsJoin returns true if the other provided type is the
// Join type or extends from the Join type.
func IsOrExtendsActivityStreamsJoin(other vocab.Type) bool {
	return typejoin.IsOrExtendsJoin(other)
}

// IsOrExtendsActivityStreamsLeave returns true if the other provided type is the
// Leave type or extends from the Leave type.
func IsOrExtendsActivityStreamsLeave(other vocab.Type) bool {
	return typeleave.IsOrExtendsLeave(other)
}

// IsOrExtendsActivityStreamsLike returns true if the other provided type is the
// Like type or extends from the Like type.
func IsOrExtendsActivityStreamsLike(other vocab.Type) bool {
	return typelike.IsOrExtendsLike(other)
}

// IsOrExtendsActivityStreamsLink returns true if the other provided type is the
// Link type or extends from the Link type.
func IsOrExtendsActivityStreamsLink(other vocab.Type) bool {
	return typelink.IsOrExtendsLink(other)
}

// IsOrExtendsActivityStreamsListen returns true if the other provided type is the
// Listen type or extends from the Listen type.
func IsOrExtendsActivityStreamsListen(other vocab.Type) bool {
	return typelisten.IsOrExtendsListen(other)
}

// IsOrExtendsActivityStreamsMention returns true if the other provided type is
// the Mention type or extends from the Mention type.
func IsOrExtendsActivityStreamsMention(other vocab.Type) bool {
	return typemention.IsOrExtendsMention(other)
}

// IsOrExtendsActivityStreamsMove returns true if the other provided type is the
// Move type or extends from the Move type.
func IsOrExtendsActivityStreamsMove(other vocab.Type) bool {
	return typemove.IsOrExtendsMove(other)
}

// IsOrExtendsActivityStreamsNote returns true if the other provided type is the
// Note type or extends from the Note type.
func IsOrExtendsActivityStreamsNote(other vocab.Type) bool {
	return typenote.IsOrExtendsNote(other)
}

// IsOrExtendsActivityStreamsObject returns true if the other provided type is the
// Object type or extends from the Object type.
func IsOrExtendsActivityStreamsObject(other vocab.Type) bool {
	return typeobject.IsOrExtendsObject(other)
}

// IsOrExtendsActivityStreamsOffer returns true if the other provided type is the
// Offer type or extends from the Offer type.
func IsOrExtendsActivityStreamsOffer(other vocab.Type) bool {
	return typeoffer.IsOrExtendsOffer(other)
}

// IsOrExtendsActivityStreamsOrderedCollection returns true if the other provided
// type is the OrderedCollection type or extends from the OrderedCollection
// type.
func IsOrExtendsActivityStreamsOrderedCollection(other vocab.Type) bool {
	return typeorderedcollection.IsOrExtendsOrderedCollection(other)
}

// IsOrExtendsActivityStreamsOrderedCollectionPage returns true if the other
// provided type is the OrderedCollectionPage type or extends from the
// OrderedCollectionPage type.
func IsOrExtendsActivityStreamsOrderedCollectionPage(other vocab.Type) bool {
	return typeorderedcollectionpage.IsOrExtendsOrderedCollectionPage(other)
}

// IsOrExtendsActivityStreamsOrganization returns true if the other provided type
// is the Organization type or extends from the Organization type.
func IsOrExtendsActivityStreamsOrganization(other vocab.Type) bool {
	return typeorganization.IsOrExtendsOrganization(other)
}

// IsOrExtendsActivityStreamsPage returns true if the other provided type is the
// Page type or extends from the Page type.
func IsOrExtendsActivityStreamsPage(other vocab.Type) bool {
	return typepage.IsOrExtendsPage(other)
}

// IsOrExtendsActivityStreamsPerson returns true if the other provided type is the
// Person type or extends from the Person type.
func IsOrExtendsActivityStreamsPerson(other vocab.Type) bool {
	return typeperson.IsOrExtendsPerson(other)
}

// IsOrExtendsActivityStreamsPlace returns true if the other provided type is the
// Place type or extends from the Place type.
func IsOrExtendsActivityStreamsPlace(other vocab.Type) bool {
	return typeplace.IsOrExtendsPlace(other)
}

// IsOrExtendsActivityStreamsProfile returns true if the other provided type is
// the Profile type or extends from the Profile type.
func IsOrExtendsActivityStreamsProfile(other vocab.Type) bool {
	return typeprofile.IsOrExtendsProfile(other)
}

// IsOrExtendsActivityStreamsPublicKey returns true if the other provided type is
// the PublicKey type or extends from the PublicKey type.
func IsOrExtendsActivityStreamsPublicKey(other vocab.Type) bool {
	return typepublickey.IsOrExtendsPublicKey(other)
}

// IsOrExtendsActivityStreamsQuestion returns true if the other provided type is
// the Question type or extends from the Question type.
func IsOrExtendsActivityStreamsQuestion(other vocab.Type) bool {
	return typequestion.IsOrExtendsQuestion(other)
}

// IsOrExtendsActivityStreamsRead returns true if the other provided type is the
// Read type or extends from the Read type.
func IsOrExtendsActivityStreamsRead(other vocab.Type) bool {
	return typeread.IsOrExtendsRead(other)
}

// IsOrExtendsActivityStreamsReject returns true if the other provided type is the
// Reject type or extends from the Reject type.
func IsOrExtendsActivityStreamsReject(other vocab.Type) bool {
	return typereject.IsOrExtendsReject(other)
}

// IsOrExtendsActivityStreamsRelationship returns true if the other provided type
// is the Relationship type or extends from the Relationship type.
func IsOrExtendsActivityStreamsRelationship(other vocab.Type) bool {
	return typerelationship.IsOrExtendsRelationship(other)
}

// IsOrExtendsActivityStreamsRemove returns true if the other provided type is the
// Remove type or extends from the Remove type.
func IsOrExtendsActivityStreamsRemove(other vocab.Type) bool {
	return typeremove.IsOrExtendsRemove(other)
}

// IsOrExtendsActivityStreamsService returns true if the other provided type is
// the Service type or extends from the Service type.
func IsOrExtendsActivityStreamsService(other vocab.Type) bool {
	return typeservice.IsOrExtendsService(other)
}

// IsOrExtendsActivityStreamsTentativeAccept returns true if the other provided
// type is the TentativeAccept type or extends from the TentativeAccept type.
func IsOrExtendsActivityStreamsTentativeAccept(other vocab.Type) bool {
	return typetentativeaccept.IsOrExtendsTentativeAccept(other)
}

// IsOrExtendsActivityStreamsTentativeReject returns true if the other provided
// type is the TentativeReject type or extends from the TentativeReject type.
func IsOrExtendsActivityStreamsTentativeReject(other vocab.Type) bool {
	return typetentativereject.IsOrExtendsTentativeReject(other)
}

// IsOrExtendsActivityStreamsTombstone returns true if the other provided type is
// the Tombstone type or extends from the Tombstone type.
func IsOrExtendsActivityStreamsTombstone(other vocab.Type) bool {
	return typetombstone.IsOrExtendsTombstone(other)
}

// IsOrExtendsActivityStreamsTravel returns true if the other provided type is the
// Travel type or extends from the Travel type.
func IsOrExtendsActivityStreamsTravel(other vocab.Type) bool {
	return typetravel.IsOrExtendsTravel(other)
}

// IsOrExtendsActivityStreamsUndo returns true if the other provided type is the
// Undo type or extends from the Undo type.
func IsOrExtendsActivityStreamsUndo(other vocab.Type) bool {
	return typeundo.IsOrExtendsUndo(other)
}

// IsOrExtendsActivityStreamsUpdate returns true if the other provided type is the
// Update type or extends from the Update type.
func IsOrExtendsActivityStreamsUpdate(other vocab.Type) bool {
	return typeupdate.IsOrExtendsUpdate(other)
}

// IsOrExtendsActivityStreamsVideo returns true if the other provided type is the
// Video type or extends from the Video type.
func IsOrExtendsActivityStreamsVideo(other vocab.Type) bool {
	return typevideo.IsOrExtendsVideo(other)
}

// IsOrExtendsActivityStreamsView returns true if the other provided type is the
// View type or extends from the View type.
func IsOrExtendsActivityStreamsView(other vocab.Type) bool {
	return typeview.IsOrExtendsView(other)
}