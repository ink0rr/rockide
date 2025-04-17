package molang

import (
	"regexp"
	"strings"

	"github.com/ink0rr/rockide/internal/sliceutil"
)

type Method struct {
	Name        string
	Signature   Signature
	Description string
	Deprecated  bool
}

var Prefixes = []string{"context", "math", "query", "temp", "variable"}

var molangQueries = []Method{
	{
		Name:        "above_top_solid",
		Signature:   "(x: number, z: number): number",
		Description: "Returns the height of the block immediately above the highest solid block at the input (x,z) position",
	},
	{
		Name:        "actor_count",
		Signature:   ": number",
		Description: "Returns the number of actors rendered in the last frame.",
	},
	{
		Name:        "all",
		Signature:   "(arg0: boolean, arg1: boolean, ...argv: boolean[]): boolean",
		Description: "Requires at least 3 arguments. Evaluates the first argument, then returns 1.0 if all of the following arguments evaluate to the same value as the first. Otherwise it returns 0.0.",
	},
	{
		Name:        "all_animations_finished",
		Signature:   ": boolean",
		Description: "Only valid in an animation controller. Returns 1.0 if all animations in the current animation controller state have played through at least once, else it returns 0.0.",
	},
	{
		Name:        "all_tags",
		Signature:   "(...tags: BlockAndItemTag[]): boolean",
		Description: "Returns if the item or block has all of the tags specified.",
	},
	{
		Name:        "anger_level",
		Signature:   ": number",
		Description: "Returns the anger level of the actor (0-n). On errors or if the actor has no anger level, returns 0. Available on the Server only.",
	},
	{
		Name:        "anim_time",
		Signature:   ": number",
		Description: "Returns the time in seconds since the current animation started, else 0.0 if not called within an animation.",
	},
	{
		Name:        "any",
		Signature:   "(arg0: boolean, arg1: boolean, ...argv: boolean[]): boolean",
		Description: "Requires at least 3 arguments. Evaluates the first argument, then returns 1.0 if any of the following arguments evaluate to the same value as the first. Otherwise it returns 0.0.",
	},
	{
		Name:        "any_animation_finished",
		Signature:   ": boolean",
		Description: "Only valid in an animation controller. Returns 1.0 if any animation in the current animation controller state has played through at least once, else it returns 0.0.",
	},
	{
		Name:        "any_tag",
		Signature:   "(...tags: BlockAndItemTag[]): boolean",
		Description: "Returns if the item or block has any of the tags specified.",
	},
	{
		Name:        "approx_eq",
		Signature:   "(...argv: number[]): boolean",
		Description: "Returns 1.0 if all of the arguments are within 0.000000 of each other, else 0.0.",
	},
	{
		Name:        "armor_color_slot",
		Signature:   "(slotIndex: number): unknown",
		Description: "Takes the armor slot index as a parameter, and returns the color of the armor in the requested slot. The valid values for the armor slot index are 0 (head), 1 (chest), 2 (legs), 3 (feet) and 4 (body).",
	},
	{
		Name:        "armor_damage_slot",
		Signature:   "(slotIndex: number): number",
		Description: "Takes the armor slot index as a parameter, and returns the damage value of the requested slot. The valid values for the armor slot index are 0 (head), 1 (chest), 2 (legs), 3 (feet) and 4 (body). Support for entities other than players may be limited, as the damage value is not always available on clients.",
	},
	{
		Name:        "armor_material_slot",
		Signature:   "(slotIndex: number): unknown",
		Description: "Takes the armor slot index as a parameter, and returns the armor material type in the requested armor slot. The valid values for the armor slot index are 0 (head), 1 (chest), 2 (legs) and 3 (feet).",
	},
	{
		Name:        "armor_texture_slot",
		Signature:   "(slotIndex: number): unknown",
		Description: "Takes the armor slot index as a parameter, and returns the texture type of the requested slot. The valid values for the armor slot index are 0 (head), 1 (chest), 2 (legs), 3 (feet) and 4 (body).",
	},
	{
		Name:        "average_frame_time",
		Signature:   "(n?: number): number",
		Description: "Returns the time in *seconds* of the average frame time over the last 'n' frames. If an argument is passed, it is assumed to be the number of frames in the past that you wish to query. 'query.average_frame_time' (or the equivalent 'query.average_frame_time(0)') will return the frame time of the frame before the current one. 'query.average_frame_time(1)' will return the average frame time of the previous two frames. Currently we store the history of the last 30 frames, although note that this may change in the future. Asking for more frames will result in only sampling the number of frames stored.",
	},
	{
		Name:        "block_face",
		Signature:   ": number",
		Description: "Returns the block face for this (only valid for certain triggers such as placing blocks, or interacting with block) (Down=0.0, Up=1.0, North=2.0, South=3.0, West=4.0, East=5.0, Undefined=6.0).",
	},
	{
		Name:        "block_has_all_tags",
		Signature:   "(origin: vector3, ...tags: BlockTag[]): boolean",
		Description: "Takes a world-origin-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has all of the tags provided.",
	},
	{
		Name:        "block_has_any_tag",
		Signature:   "(origin: vector3, ...tags: BlockTag[]): boolean",
		Description: "Takes a world-origin-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has any of the tags provided.",
	},
	{
		Name:        "block_neighbor_has_all_tags",
		Signature:   "(origin: vector3, ...tags: BlockTag[]): boolean",
		Description: "Takes a block-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has all of the tags provided.",
	},
	{
		Name:        "block_neighbor_has_any_tag",
		Signature:   "(origin: vector3, ...tags: BlockTag[]): boolean",
		Description: "Takes a block-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has any of the tags provided.",
	},
	{
		Name:        "block_property",
		Signature:   "(identifier: string): string | number | boolean",
		Description: "(No longer available in pack min_engine_version 1.20.40.) Returns the value of the associated block's Block State.",
		Deprecated:  true,
	},
	{
		Name:        "block_state",
		Signature:   ": string | number | boolean",
		Description: "Returns the value of the associated block's Block State.",
	},
	{
		Name:        "blocking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is blocking, else it returns 0.0.",
	},
	{
		Name:        "body_x_rotation",
		Signature:   ": number",
		Description: "Returns the body pitch rotation if called on an actor, else it returns 0.0.",
	},
	{
		Name:        "body_y_rotation",
		Signature:   ": number",
		Description: "Returns the body yaw rotation if called on an actor, else it returns 0.0.",
	},
	{
		Name:        "bone_aabb",
		Signature:   "(bone: string): { min: vector3, max: vector3 }",
		Description: "Returns the axis aligned bounding box of a bone as a struct with members '.min', '.max', along with '.x', '.y', and '.z' values for each.",
	},
	{
		Name:        "bone_orientation_matrix",
		Signature:   "(bone: string): unknown",
		Description: "Takes the name of the bone as an argument. Returns the bone orientation (as a matrix) of the desired bone provided it exists in the queryable geometry of the mob, else this returns the identity matrix and throws a content error.",
	},
	{
		Name:        "bone_orientation_trs",
		Signature:   "(bone: string): { t: vector3, r: vector3, s: vector3}",
		Description: "TRS stands for Translate/Rotate/Scale. Takes the name of the bone as an argument. Returns the bone orientation matrix decomposed into the component translation/rotation/scale parts of the desired bone provided it exists in the queryable geometry of the mob, else this returns the identity matrix and throws a content error. The returned value is returned as a variable of type 'struct' with members '.t', '.r', and '.s', each with members '.x', '.y', and '.z', and can be accessed as per the following example: v.my_variable = q.bone_orientation_trs('rightarm'); return v.my_variable.r.x;",
	},
	{
		Name:        "bone_origin",
		Signature:   ": vector3",
		Description: "Returns the initial (from the .geo) pivot of a bone as a struct with members '.x', '.y', and '.z'.",
	},
	{
		Name:        "bone_rotation",
		Signature:   ": vector3",
		Description: "Returns the initial (from the .geo) rotation of a bone as a struct with members '.x', '.y', and '.z' in degrees.",
	},
	{
		Name:        "camera_distance_range_lerp",
		Signature:   "(a: number, b: number): number",
		Description: "Takes two distances (any order) and return a number from 0 to 1 based on the camera distance between the two ranges clamped to that range. For example, 'query.camera_distance_range_lerp(10, 20)' will return 0 for any distance less than or equal to 10, 0.2 for a distance of 12, 0.5 for 15, and 1 for 20 or greater. If you pass in (20, 10), a distance of 20 will return 0.0.",
	},
	{
		Name:        "camera_rotation",
		Signature:   "(axis: number): number",
		Description: "Returns the rotation of the camera. Requires one argument representing the rotation axis you would like (0 for x, 1 for y).",
	},
	{
		Name:        "can_climb",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can climb, else it returns 0.0.",
	},
	{
		Name:        "can_damage_nearby_mobs",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can damage nearby mobs, else it returns 0.0.",
	},
	{
		Name:        "can_dash",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can dash, else it returns 0.0",
	},
	{
		Name:        "can_fly",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can fly, else it returns 0.0.",
	},
	{
		Name:        "can_power_jump",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can power jump, else it returns 0.0.",
	},
	{
		Name:        "can_swim",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can swim, else it returns 0.0.",
	},
	{
		Name:        "can_walk",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity can walk, else it returns 0.0.",
	},
	{
		Name:        "cape_flap_amount",
		Signature:   ": number",
		Description: "Returns value between 0.0 and 1.0 with 0.0 meaning cape is fully down and 1.0 is cape is fully up.",
	},
	{
		Name:        "cardinal_block_face_placed_on",
		Signature:   ": number",
		Description: "DEPRECATED (please use query.block_face instead) Returns the block face for this (only valid for on_placed_by_player trigger) (Down=0.0, Up=1.0, North=2.0, South=3.0, West=4.0, East=5.0, Undefined=6.0).",
		Deprecated:  true,
	},
	{
		Name:        "cardinal_facing",
		Signature:   ": number",
		Description: "Returns the current facing of the player (Down=0.0, Up=1.0, North=2.0, South=3.0, West=4.0, East=5.0, Undefined=6.0).",
	},
	{
		Name:        "cardinal_facing_2d",
		Signature:   ": number",
		Description: "Returns the current facing of the player ignoring up/down part of the direction (North=2.0, South=3.0, West=4.0, East=5.0, Undefined=6.0).",
	},
	{
		Name:        "cardinal_player_facing",
		Signature:   ": number",
		Description: "Returns the current facing of the player (Down=0.0, Up=1.0, North=2.0, South=3.0, West=4.0, East=5.0, Undefined=6.0).",
	},
	{
		Name:        "client_max_render_distance",
		Signature:   ": number",
		Description: "Returns the max render distance in chunks of the current client. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "client_memory_tier",
		Signature:   ": number",
		Description: "Returns a number representing the client RAM memory tier, 0 = 'SuperLow', 1 = 'Low', 2 = 'Mid', 3 = 'High', or 4 = 'SuperHigh'. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "combine_entities",
		Signature:   "(...refs: any[]): entity[]", // TODO: undocumented
		Description: "Combines any valid entity references from all arguments into a single array. Note that order is not preserved, and duplicates and invalid values are removed.",
	},
	{
		Name:        "cooldown_time",
		Signature:   "(slotName: EquipmentSlot, slotIndex?: number): number",
		Description: "Returns the total cooldown time in seconds for the item held or worn by the specified equipment slot name (and if required second numerical slot id), otherwise returns 0. Uses the same name and id that the replaceitem command takes when querying entities.",
	},
	{
		Name:        "cooldown_time_remaining",
		Signature:   "(slotName: EquipmentSlot, slotIndex?: number): number",
		Description: "Returns the cooldown time remaining in seconds for specified cooldown type or the item held or worn by the specified equipment slot name (and if required second numerical slot id), otherwise returns 0. Uses the same name and id that the replaceitem command takes when querying entities. Returns highest cooldown if no parameters are supplied.",
	},
	{
		Name:        "count",
		Signature:   "(arr: any[]): number",
		Description: "Counts the number of things passed to it (arrays are counted as the number of elements they contain; non-arrays count as 1).",
	},
	{
		Name:        "current_squish_value",
		Signature:   ": number",
		Description: "Returns the squish value for the current entity, or 0.0 if this doesn't make sense.",
	},
	{
		Name:        "dash_cooldown_progress",
		Signature:   ": number",
		Description: "(No longer available in pack min_engine_version 1.20.50.) DEPRECATED. DO NOT USE AFTER 1.20.40. Please see camel.entity.json script.pre_animation for example of how to now process dash cooldown. Returns dash cooldown progress if the entity can dash, else it returns 0.0.",
		Deprecated:  false,
	},
	{
		Name:        "day",
		Signature:   ": number",
		Description: "Returns the day of the current level.",
	},
	{
		Name:        "death_ticks",
		Signature:   ": number",
		Description: "Returns the elapsed ticks since the mob started dying.",
	},
	{
		Name:        "debug_output",
		Signature:   ": void",
		Description: "debug log a value to the output debug window for builds that have one",
	},
	{
		Name:        "delta_time",
		Signature:   ": number",
		Description: "Returns the time in seconds since the previous frame.",
	},
	{
		Name:        "distance_from_camera",
		Signature:   ": number",
		Description: "Returns the distance of the root of this actor or particle emitter from the camera.",
	},
	{
		Name:        "effect_emitter_count",
		Signature:   ": number",
		Description: "Returns the total number of active emitters of the callee's particle effect type.",
	},
	{
		Name:        "effect_particle_count",
		Signature:   ": number",
		Description: "Returns the total number of active particles of the callee's particle effect type.",
	},
	{
		Name:        "equipment_count",
		Signature:   ": number",
		Description: "Returns the number of equipped armor pieces for an actor from 0 to 5, not counting items held in hands. (To query for hand slots, use query.is_item_equipped or query.is_item_name_any).",
	},
	{
		Name:        "equipped_item_all_tags",
		Signature:   "(slotName: EquipmentSlot, ...tags: ItemTag[]): boolean",
		Description: "Takes a slot name followed by any tag you want to check for in the form of 'tag_name' and returns 1 if all of the tags are on that equipped item, 0 otherwise.",
	},
	{
		Name:        "equipped_item_any_tag",
		Signature:   "(slotName: EquipmentSlot, ...tags: ItemTag[]): boolean",
		Description: "Takes a slot name followed by any tag you want to check for in the form of 'tag_name' and returns 0 if none of the tags are on that equipped item or 1 if at least 1 tag exists.",
	},
	{
		Name:        "equipped_item_is_attachable",
		Signature:   "(slot: number): boolean",
		Description: "Takes the desired hand slot as a parameter (0 or 'main_hand' for main hand, 1 or 'off_hand' for off hand), and returns whether the item is an attachable or not.",
	},
	{
		Name:        "eye_target_x_rotation",
		Signature:   ": number",
		Description: "Returns the X eye rotation of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "eye_target_y_rotation",
		Signature:   ": number",
		Description: "Returns the Y eye rotation of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "facing_target_to_range_attack",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is attacking from range (i.e. minecraft:behavior.ranged_attack), else it returns 0.0.",
	},
	{
		Name:        "frame_alpha",
		Signature:   ": number",
		Description: "Returns the ratio (from 0 to 1) of how much between AI ticks this frame is being rendered.",
	},
	{
		Name:        "get_actor_info_id",
		Signature:   ": string",
		Description: "Returns the integer id of an actor by its string name.",
	},
	{
		Name:        "get_animation_frame",
		Signature:   ": number",
		Description: "Returns the current texture of the item",
	},
	{
		Name:        "get_default_bone_pivot",
		Signature:   "(bone: string, orientation: number): number",
		Description: "Gets specified axis of the specified bone orientation pivot.",
	},
	{
		Name:        "get_equipped_item_name",
		Signature:   ": string",
		Description: "DEPRECATED (Use query.is_item_name_any instead if possible so names can be changed later without breaking content.) Takes one optional hand slot as a parameter (0 or 'main_hand' for main hand, 1 or 'off_hand' for off hand), and a second parameter (0=default) if you would like the equipped item or any non-zero number for the currently rendered item, and returns the name of the item in the requested slot (defaulting to the main hand if no parameter is supplied) if there is one, otherwise returns ''.",
		Deprecated:  true,
	},
	{
		Name:        "get_locator_offset",
		Signature:   ": unknown", // TODO: undocumented
		Description: "Gets specified axis of the specified locator offset.",
	},
	{
		Name:        "get_name",
		Signature:   ": string",
		Description: "DEPRECATED (Use query.is_name_any instead if possible so names can be changed later without breaking content.)Get the name of the mob if there is one, otherwise return ''.",
		Deprecated:  true,
	},
	{
		Name:        "get_root_locator_offset",
		Signature:   ": unknown", // TODO: undocumented
		Description: "Gets specified axis of the specified locator offset of the root model.",
	},
	{
		Name:        "ground_speed",
		Signature:   ": number",
		Description: "Returns the ground speed of the entity in meters/second.",
	},
	{
		Name:        "had_component_group",
		Signature:   "(componentGroup: string): boolean",
		Description: "Usable only in behavior packs when determining the default value for an entity's Property. Requires one string argument. If the entity is being loaded from data that was last saved with a component_group with the specified name, returns 1.0, otherwise returns 0.0. The purpose of this query is to allow entity definitions to change and still be able to load the correct state of entities.",
	},
	{
		Name:        "has_any_family",
		Signature:   "(...families: TypeFamily[]): boolean",
		Description: "Returns 1 if the entity has any of the specified families, else 0.",
	},
	{
		Name:        "has_armor_slot",
		Signature:   "(slotIndex: number): boolean",
		Description: "Takes the armor slot index as a parameter, and returns 1.0 if the entity has armor in the requested slot, else it returns 0.0. The valid values for the armor slot index are 0 (head), 1 (chest), 2 (legs) and 3 (feet).",
	},
	{
		Name:        "has_biome_tag",
		Signature:   "(tag: BiomeTag): boolean",
		Description: "Returns whether or not a Block Placement Target has a specific biome tag",
	},
	{
		Name:        "has_block_property",
		Signature:   "(identifier: string): boolean",
		Description: "(No longer available in pack min_engine_version 1.20.40.) Returns 1.0 if the associated block has the given block state or 0.0 if not.",
		Deprecated:  true,
	},
	{
		Name:        "has_block_state",
		Signature:   "(name: string): boolean",
		Description: "Returns 1.0 if the associated block has the given block state or 0.0 if not.",
	},
	{
		Name:        "has_cape",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the player has a cape, else it returns 0.0.",
	},
	{
		Name:        "has_collision",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has collisions enabled, else it returns 0.0.",
	},
	{
		Name:        "has_dash_cooldown",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has cooldown on its dash, else it returns 0.0",
	},
	{
		Name:        "has_gravity",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is affected by gravity, else it returns 0.0.",
	},
	{
		Name:        "has_head_gear",
		Signature:   ": boolean",
		Description: "Returns boolean whether an Actor has an item in their head armor slot or not, or false if no actor in current context",
	},
	{
		Name:        "has_owner",
		Signature:   ": boolean",
		Description: "Returns true if the entity has an owner ID else it returns false",
	},
	{
		Name:        "has_player_rider",
		Signature:   ": boolean",
		Description: "Returns 1 if the entity has a player riding it in any seat, else it returns 0.",
	},
	{
		Name:        "has_property",
		Signature:   "(identifier: EntityProperty): boolean",
		Description: "Takes one argument: the name of the property on the Actor. Returns 1.0 if a property with the given name exists, 0 otherwise.",
	},
	{
		Name:        "has_rider",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has a rider, else it returns 0.0",
	},
	{
		Name:        "has_target",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has a target, else it returns 0.0",
	},
	{
		Name:        "head_roll_angle",
		Signature:   ": number",
		Description: "Returns the roll angle of the head of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "head_x_rotation",
		Signature:   "(n: number): number",
		Description: "Takes one argument as a parameter. Returns the nth head x rotation of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "head_y_rotation",
		Signature:   "(n: number, maxRotation?: number): number",
		Description: "Takes one argument as a parameter. Returns the nth head y rotation of the entity if it makes sense, else it returns 0.0. Horses, zombie horses, skeleton horses, donkeys and mules require a second parameter that clamps rotation in degrees.",
	},
	{
		Name:        "health",
		Signature:   ": number",
		Description: "Returns the health of the entity, or 0.0 if it doesn't make sense to call on this entity.",
	},
	{
		Name:        "heartbeat_interval",
		Signature:   ": number",
		Description: "Returns the heartbeat interval of the actor in seconds. Returns 0 when the actor has no heartbeat.",
	},
	{
		Name:        "heartbeat_phase",
		Signature:   ": number",
		Description: "Returns the heartbeat phase of the actor. 0.0 if at start of current heartbeat, 1.0 if at the end. Returns 0 on errors or when the actor has no heartbeat. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "heightmap",
		Signature:   "(x: number, z: number): number", // TODO: undocumented
		Description: "Queries Height Map",
	},
	{
		Name:        "hurt_direction",
		Signature:   ": number",
		Description: "Returns the hurt direction for the actor, otherwise returns 0.",
	},
	{
		Name:        "hurt_time",
		Signature:   ": number",
		Description: "Returns the hurt time for the actor, otherwise returns 0.",
	},
	{
		Name:        "in_range",
		Signature:   "(value: number, min: number, max: number): boolean",
		Description: "Requires 3 numerical arguments: some value, a minimum, and a maximum. If the first argument is between the minimum and maximum (inclusive), returns 1.0. Otherwise returns 0.0.",
	},
	{
		Name:        "invulnerable_ticks",
		Signature:   ": number",
		Description: "Returns the number of ticks of invulnerability the entity has left if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "is_admiring",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is admiring, else it returns 0.0.",
	},
	{
		Name:        "is_alive",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is alive, and 0.0 if it's dead.",
	},
	{
		Name:        "is_angry",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is angry, else it returns 0.0.",
	},
	{
		Name:        "is_attached",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is attached to another entity (such as being held or worn), else it will return 0.0. Available only with resource packs.",
	},
	{
		Name:        "is_attached_to_entity",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the actor is attached to an entity, else it will return 0.0.",
	},
	{
		Name:        "is_avoiding_block",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is fleeing from a block, else it returns 0.0.",
	},
	{
		Name:        "is_avoiding_mobs",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is fleeing from mobs, else it returns 0.0.",
	},
	{
		Name:        "is_baby",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is a baby, else it returns 0.0.",
	},
	{
		Name:        "is_breathing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is breathing, else it returns 0.0.",
	},
	{
		Name:        "is_bribed",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has been bribed, else it returns 0.0.",
	},
	{
		Name:        "is_carrying_block",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is carrying a block, else it returns 0.0.",
	},
	{
		Name:        "is_casting",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is casting, else it returns 0.0.",
	},
	{
		Name:        "is_celebrating",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is celebrating, else it returns 0.0.",
	},
	{
		Name:        "is_celebrating_special",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is doing a special celebration, else it returns 0.0.",
	},
	{
		Name:        "is_charged",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is charged, else it returns 0.0.",
	},
	{
		Name:        "is_charging",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is charging, else it returns 0.0.",
	},
	{
		Name:        "is_chested",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has chests attached to it, else it returns 0.0.",
	},
	{
		Name:        "is_cooldown_type",
		Signature:   "(cooldownType: string, slotName: EquipmentSlot, slotIndex?: number): boolean",
		Description: "Returns 1.0 if the specified held or worn item has the specified cooldown type name, otherwise returns 0.0. First argument is the cooldown name to check for, second argument is the equipment slot name, and if required third argument is the numerical slot id. For second and third arguments, uses the same name and id that the replaceitem command takes when querying entities.",
	},
	{
		Name:        "is_crawling",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is crawling, else it returns 0.0",
	},
	{
		Name:        "is_critical",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is critical, else it returns 0.0.",
	},
	{
		Name:        "is_croaking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is croaking, else it returns 0.0.",
	},
	{
		Name:        "is_dancing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is dancing, else it returns 0.0.",
	},
	{
		Name:        "is_delayed_attacking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is attacking using the delayed attack, else it returns 0.0.",
	},
	{
		Name:        "is_digging",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is digging, else it returns 0.0.",
	},
	{
		Name:        "is_eating",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is eating, else it returns 0.0.",
	},
	{
		Name:        "is_eating_mob",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is eating a mob, else it returns 0.0.",
	},
	{
		Name:        "is_elder",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is an elder version of it, else it returns 0.0.",
	},
	{
		Name:        "is_emerging",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is emerging, else it returns 0.0.",
	},
	{
		Name:        "is_emoting",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is emoting, else it returns 0.0.",
	},
	{
		Name:        "is_enchanted",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is enchanted, else it returns 0.0.",
	},
	{
		Name:        "is_feeling_happy",
		Signature:   ": boolean",
		Description: "(No longer available in pack min_engine_version 1.20.50.) DEPRECATED after 1.20.40. Returns 1.0 if behavior.timer_flag_2 is running, else it returns 0.0.",
		Deprecated:  true,
	},
	{
		Name:        "is_fire_immune",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is immune to fire, else it returns 0.0.",
	},
	{
		Name:        "is_first_person",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is being rendered in first person mode, else it returns 0.0.",
	},
	{
		Name:        "is_ghost",
		Signature:   ": boolean",
		Description: "Returns 1.0 if an entity is a ghost, else it returns 0.0.",
	},
	{
		Name:        "is_gliding",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is gliding, else it returns 0.0.",
	},
	{
		Name:        "is_grazing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is grazing, or 0.0 if not.",
	},
	{
		Name:        "is_idling",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is idling, else it returns 0.0.",
	},
	{
		Name:        "is_ignited",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is ignited, else it returns 0.0.",
	},
	{
		Name:        "is_illager_captain",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is an illager captain, else it returns 0.0.",
	},
	{
		Name:        "is_in_contact_with_water",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is in contact with any water (water, rain, splash water bottle), else it returns 0.0.",
	},
	{
		Name:        "is_in_lava",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is in lava, else it returns 0.0.",
	},
	{
		Name:        "is_in_love",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is in love, else it returns 0.0.",
	},
	{
		Name:        "is_in_ui",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is rendered as part of the UI, else it returns 0.0.",
	},
	{
		Name:        "is_in_water",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is in water, else it returns 0.0.",
	},
	{
		Name:        "is_in_water_or_rain",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is in water or rain, else it returns 0.0.",
	},
	{
		Name:        "is_interested",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is interested, else it returns 0.0.",
	},
	{
		Name:        "is_invisible",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is invisible, else it returns 0.0.",
	},
	{
		Name:        "is_item_equipped",
		Signature:   "(slot?: number): boolean",
		Description: "Takes one optional hand slot as a parameter (0 or 'main_hand' for main hand, 1 or 'off_hand' for off hand), and returns 1.0 if there is an item in the requested slot (defaulting to the main hand if no parameter is supplied), otherwise returns 0.0.",
	},
	{
		Name:        "is_item_name_any",
		Signature:   "(slotName: EquipmentSlot, slotIndex: number, ...identifiers: ItemIdentifier[]): boolean",
		Description: "Takes an equipment slot name (see the replaceitem command) and an optional slot index value. (The slot index is required for slot names that have multiple slots, for example 'slot.hotbar'.) After that, takes one or more full name (with 'namespace:') strings to check for. Returns 1.0 if an item in the specified slot has any of the specified names, otherwise returns 0.0. An empty string '' can be specified to check for an empty slot. Note that querying slot.enderchest, slot.saddle, slot.armor, or slot.chest will only work in behavior packs. A preferred query to query.get_equipped_item_name, as it can be adjusted by Mojang to avoid breaking content if names are changed.",
	},
	{
		Name:        "is_jump_goal_jumping",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is doing a jump goal jump, else it returns 0.0.",
	},
	{
		Name:        "is_jumping",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is jumping, else it returns 0.0.",
	},
	{
		Name:        "is_laying_down",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is laying down, else it returns 0.0.",
	},
	{
		Name:        "is_laying_egg",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is laying an egg, else it returns 0.0.",
	},
	{
		Name:        "is_leashed",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is leashed to something, else it returns 0.0.",
	},
	{
		Name:        "is_levitating",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is levitating, else it returns 0.0.",
	},
	{
		Name:        "is_lingering",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is lingering, else it returns 0.0.",
	},
	{
		Name:        "is_local_player",
		Signature:   ": boolean",
		Description: "Takes no arguments. Returns 1.0 if the entity is the local player for the current game window, else it returns 0.0. In splitscreen returns 0.0 for the other local players for other views. Always returns 0.0 if used in a behavior pack.",
	},
	{
		Name:        "is_moving",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is moving, else it returns 0.0.",
	},
	{
		Name:        "is_name_any",
		Signature:   "(...names: string[]): boolean",
		Description: "Takes one or more arguments. If the entity's name is any of the specified string values, returns 1.0. Otherwise returns 0.0. A preferred query to query.get_name, as it can be adjusted by Mojang to avoid breaking content if names are changed.",
	},
	{
		Name:        "is_on_fire",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is on fire, else it returns 0.0.",
	},
	{
		Name:        "is_on_ground",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is on the ground, else it returns 0.0.",
	},
	{
		Name:        "is_on_screen",
		Signature:   ": boolean",
		Description: "Returns 1.0 if this is called on an entity at a time when it is known if it is on screen, else it returns 0.0.",
	},
	{
		Name:        "is_onfire",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is on fire, else it returns 0.0.",
	},
	{
		Name:        "is_orphaned",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is orphaned, else it returns 0.0.",
	},
	{
		Name:        "is_owner_identifier_any",
		Signature:   "(...identifiers: EntityIdentifier[]): boolean",
		Description: "Takes one or more arguments. Returns whether the root actor identifier is any of the specified strings. A preferred query to query.owner_identifier, as it can be adjusted by Mojang to avoid breaking content if names are changed.",
	},
	{
		Name:        "is_persona_or_premium_skin",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the player has a persona or premium skin, else it returns 0.0.",
	},
	{
		Name:        "is_playing_dead",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is playing dead, else it returns 0.0.",
	},
	{
		Name:        "is_powered",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is powered, else it returns 0.0.",
	},
	{
		Name:        "is_pregnant",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is pregnant, else it returns 0.0.",
	},
	{
		Name:        "is_ram_attacking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is using a ram attack, else it returns 0.0.",
	},
	{
		Name:        "is_resting",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is resting, else it returns 0.0.",
	},
	{
		Name:        "is_riding",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is riding, else it returns 0.0.",
	},
	{
		Name:        "is_rising",
		Signature:   ": boolean",
		Description: "(No longer available in pack min_engine_version 1.20.50.) DEPRECATED after 1.20.40. Returns 1.0 if behavior.timer_flag_2 is running, else it returns 0.0.",
		Deprecated:  true,
	},
	{
		Name:        "is_roaring",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is currently roaring, else it returns 0.0.",
	},
	{
		Name:        "is_rolling",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is rolling, else it returns 0.0.",
	},
	{
		Name:        "is_saddled",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity has a saddle, else it returns 0.0.",
	},
	{
		Name:        "is_scared",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is scared, else it returns 0.0.",
	},
	{
		Name:        "is_scenting",
		Signature:   ": boolean",
		Description: "(No longer available in pack min_engine_version 1.20.50.) DEPRECATED after 1.20.40. Returns 1.0 if behavior.timer_flag_1 is running, else it returns 0.0.",
		Deprecated:  true,
	},
	{
		Name:        "is_searching",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is searching, else it returns 0.0.",
	},
	{
		Name:        "is_selected_item",
		Signature:   ": boolean",
		Description: "Returns true if the player has selected an item in the inventory, else it returns 0.0.",
	},
	{
		Name:        "is_shaking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is casting, else it returns 0.0.",
	},
	{
		Name:        "is_shaking_wetness",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is shaking water off, else it returns 0.0.",
	},
	{
		Name:        "is_sheared",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is able to be sheared and is sheared, else it returns 0.0.",
	},
	{
		Name:        "is_shield_powered",
		Signature:   ": boolean",
		Description: "Returns 1.0f if the entity has an active powered shield if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "is_silent",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is silent, else it returns 0.0.",
	},
	{
		Name:        "is_sitting",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sitting, else it returns 0.0.",
	},
	{
		Name:        "is_sleeping",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sleeping, else it returns 0.0.",
	},
	{
		Name:        "is_sneaking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sneaking, else it returns 0.0.",
	},
	{
		Name:        "is_sneezing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sneezing, else it returns 0.0.",
	},
	{
		Name:        "is_sniffing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sniffing, else it returns 0.0.",
	},
	{
		Name:        "is_sonic_boom",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is using sonic boom, else it returns 0.0.",
	},
	{
		Name:        "is_spectator",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is spectator, else it returns 0.0.",
	},
	{
		Name:        "is_sprinting",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is sprinting, else it returns 0.0.",
	},
	{
		Name:        "is_stackable",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is stackable, else it returns 0.0.",
	},
	{
		Name:        "is_stalking",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is stalking, else it returns 0.0.",
	},
	{
		Name:        "is_standing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is standing, else it returns 0.0.",
	},
	{
		Name:        "is_stunned",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is currently stunned, else it returns 0.0.",
	},
	{
		Name:        "is_swimming",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is swimming, else it returns 0.0.",
	},
	{
		Name:        "is_tamed",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is tamed, else it returns 0.0.",
	},
	{
		Name:        "is_transforming",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is transforming, else it returns 0.0.",
	},
	{
		Name:        "is_using_item",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is using an item, else it returns 0.0.",
	},
	{
		Name:        "is_wall_climbing",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is climbing a wall, else it returns 0.0.",
	},
	{
		Name:        "item_in_use_duration",
		Signature:   ": number",
		Description: "Returns the amount of time an item has been in use in seconds up to the maximum duration, else 0.0 if it doesn't make sense.",
	},
	{
		Name:        "item_is_charged",
		Signature:   "(slot?: number): boolean",
		Description: "Takes one optional hand slot as a parameter (0 or 'main_hand' for main hand, 1 or 'off_hand' for off hand), and returns 1.0 if the item is charged in the requested slot (defaulting to the main hand if no parameter is supplied), otherwise returns 0.0.",
	},
	{
		Name:        "item_max_use_duration",
		Signature:   ": number",
		Description: "Returns the maximum amount of time the item can be used, else 0.0 if it doesn't make sense.",
	},
	{
		Name:        "item_remaining_use_duration",
		Signature:   "(slotName: EquipmentSlot): number",
		Description: "Returns the amount of time an item has left to use, else 0.0 if it doesn't make sense. Item queried is specified by the slot name 'main_hand' or 'off_hand'. Time remaining is normalized using the normalization value, only if one is given, else it is returned in seconds.",
	},
	{
		Name:        "item_slot_to_bone_name",
		Signature:   "(slotName: EquipmentSlot): string",
		Description: "query.item_slot_to_bone_name requires one parameter: the name of the equipment slot. This function returns the name of the bone this entity has mapped to that slot.",
	},
	{
		Name:        "key_frame_lerp_time",
		Signature:   ": number",
		Description: "Returns the ratio between the previous and next key frames.",
	},
	{
		Name:        "last_frame_time",
		Signature:   "(prevFrame?: number): number",
		Description: "Returns the time in *seconds* of the last frame. If an argument is passed, it is assumed to be the number of frames in the past that you wish to query. 'query.last_frame_time' (or the equivalent 'query.last_frame_time(0)') will return the frame time of the frame before the current one. 'query.last_frame_time(1)' will return the frame time of two frames ago. Currently we store the history of the last 30 frames, although note that this may change in the future. Passing an index more than the available data will return the oldest frame stored.",
	},
	{
		Name:        "last_hit_by_player",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity was last hit by the player, else it returns 0.0. If called by the client always returns 0.0.",
	},
	{
		Name:        "last_input_mode_is_any",
		Signature:   "(...inputModes: InputMode[]): boolean",
		Description: "Takes one or more arguments ('keyboard_and_mouse', 'touch', 'gamepad', or 'motion_controller'). If the last input used is any of the specified string values, returns 1.0. Otherwise returns 0.0. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "lie_amount",
		Signature:   ": number",
		Description: "Returns the lie down amount for the entity.",
	},
	{
		Name:        "life_span",
		Signature:   ": number",
		Description: "Returns the limited life span of an entity, or 0.0 if it lives forever",
	},
	{
		Name:        "life_time",
		Signature:   ": number",
		Description: "Returns the time in seconds since the current animation started, else 0.0 if not called within an animation.",
	},
	{
		Name:        "lod_index",
		Signature:   "(...distances: number[]): number",
		Description: "Takes an array of distances and returns the zero - based index of which range the actor is in based on distance from the camera. For example, 'query.lod_index(10, 20, 30)' will return 0, 1, or 2 based on whether the mob is less than 10, 20, or 30 units away from the camera, or it will return 3 if it is greater than 30.",
	},
	{
		Name:        "log",
		Signature:   ": void",
		Description: "debug log a value to the content log",
	},
	{
		Name:        "main_hand_item_max_duration",
		Signature:   ": number",
		Description: "Returns the use time maximum duration for the main hand item if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "main_hand_item_use_duration",
		Signature:   ": number",
		Description: "Returns the use time for the main hand item.",
	},
	{
		Name:        "mark_variant",
		Signature:   ": number",
		Description: "Returns the entity's mark variant",
	},
	{
		Name:        "max_durability",
		Signature:   ": number",
		Description: "Returns the max durability an item can take.",
	},
	{
		Name:        "max_health",
		Signature:   ": number",
		Description: "Returns the maximum health of the entity, or 0.0 if it doesn't make sense to call on this entity.",
	},
	{
		Name:        "max_trade_tier",
		Signature:   ": number",
		Description: "Returns the maximum trade tier of the entity if it makes sense, else it returns 0.0",
	},
	{
		Name:        "maximum_frame_time",
		Signature:   "(prevFrame?: number): number",
		Description: "Returns the time in *seconds* of the most expensive frame over the last 'n' frames. If an argument is passed, it is assumed to be the number of frames in the past that you wish to query. 'query.maximum_frame_time' (or the equivalent 'query.maximum_frame_time(0)') will return the frame time of the frame before the current one. 'query.maximum_frame_time(1)' will return the maximum frame time of the previous two frames. Currently we store the history of the last 30 frames, although note that this may change in the future. Asking for more frames will result in only sampling the number of frames stored.",
	},
	{
		Name:        "minimum_frame_time",
		Signature:   "(prevFrame?: number): number",
		Description: "Returns the time in *seconds* of the least expensive frame over the last 'n' frames. If an argument is passed, it is assumed to be the number of frames in the past that you wish to query. 'query.minimum_frame_time' (or the equivalent 'query.minimum_frame_time(0)') will return the frame time of the frame before the current one. 'query.minimum_frame_time(1)' will return the minimum frame time of the previous two frames. Currently we store the history of the last 30 frames, although note that this may change in the future. Asking for more frames will result in only sampling the number of frames stored.",
	},
	{
		Name:        "model_scale",
		Signature:   ": number",
		Description: "Returns the scale of the current entity.",
	},
	{
		Name:        "modified_distance_moved",
		Signature:   ": number",
		Description: "Returns the total distance the entity has moved horizontally in meters (since the entity was last loaded, not necessarily since it was originally created) modified along the way by status flags such as is_baby or on_fire.",
	},
	{
		Name:        "modified_move_speed",
		Signature:   ": number",
		Description: "Returns the current walk speed of the entity modified by status flags such as is_baby or on_fire.",
	},
	{
		Name:        "moon_brightness",
		Signature:   ": number",
		Description: "Returns the brightness of the moon (FULL_MOON=1.0, WANING_GIBBOUS=0.75, FIRST_QUARTER=0.5, WANING_CRESCENT=0.25, NEW_MOON=0.0, WAXING_CRESCENT=0.25, LAST_QUARTER=0.5, WAXING_GIBBOUS=0.75).",
	},
	{
		Name:        "moon_phase",
		Signature:   ": number",
		Description: "Returns the phase of the moon (FULL_MOON=0, WANING_GIBBOUS=1, FIRST_QUARTER=2, WANING_CRESCENT=3, NEW_MOON=4, WAXING_CRESCENT=5, LAST_QUARTER=6, WAXING_GIBBOUS=7).",
	},
	{
		Name:        "movement_direction",
		Signature:   ": number",
		Description: "Returns the specified axis of the normalized position delta of the entity.",
	},
	{
		Name:        "noise",
		Signature:   "(...argv: number[]): number",
		Description: "Queries Perlin Noise Map",
	},
	{
		Name:        "on_fire_time",
		Signature:   ": number",
		Description: "Returns the time that the entity is on fire, else it returns 0.0.",
	},
	{
		Name:        "out_of_control",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the entity is out of control, else it returns 0.0.",
	},
	{
		Name:        "overlay_alpha",
		Signature:   ": number",
		Description: "DEPRECATED (Do not use - this function is deprecated and will be removed).",
		Deprecated:  true,
	},
	{
		Name:        "owner_identifier",
		Signature:   ": string",
		Description: "DEPRECATED (Use query.is_owner_identifier_any instead if possible so names can be changed later without breaking content.) Returns the root actor identifier.",
		Deprecated:  true,
	},
	{
		Name:        "player_level",
		Signature:   ": number",
		Description: "Returns the players level if the actor is a player, otherwise returns 0.",
	},
	{
		Name:        "position",
		Signature:   "(axis: number): number",
		Description: "Returns the absolute position of an actor. Takes one argument that represents the desired axis (0 == x-axis, 1 == y-axis, 2 == z-axis).",
	},
	{
		Name:        "position_delta",
		Signature:   "(axis: number): number",
		Description: "Returns the position delta for an actor. Takes one argument that represents the desired axis (0 == x-axis, 1 == y-axis, 2 == z-axis).",
	},
	{
		Name:        "previous_squish_value",
		Signature:   ": number",
		Description: "Returns the previous squish value for the current entity, or 0.0 if this doesn't make sense.",
	},
	{
		Name:        "property",
		Signature:   "(identifier: EntityProperty): string | number | boolean",
		Description: "Takes one argument: the name of the property on the entity. Returns the value of that property if it exists, else 0.0 if not.",
	},
	{
		Name:        "relative_block_has_all_tags",
		Signature:   "(position: vector3, ...tags: BlockTag[]): number", // TODO: undocumented
		Description: "Takes an entity-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has all of the tags provided.",
	},
	{
		Name:        "relative_block_has_any_tag",
		Signature:   "(position: vector3, ...tags: BlockTag[]): number", // TODO: undocumented
		Description: "Takes an entity-relative position and one or more tag names, and returns either 0 or 1 based on if the block at that position has any of the tags provided.",
	},
	{
		Name:        "remaining_durability",
		Signature:   ": number",
		Description: "Returns how much durability an item has remaining.",
	},
	{
		Name:        "ride_body_x_rotation",
		Signature:   ": number",
		Description: "Returns the body pitch world-rotation of the ride an entity, else it returns 0.0.",
	},
	{
		Name:        "ride_body_y_rotation",
		Signature:   ": number",
		Description: "Returns the body yaw world-rotation of the ride of on an entity, else it returns 0.0.",
	},
	{
		Name:        "ride_head_x_rotation",
		Signature:   ": number",
		Description: "Returns the head x world-rotation of the ride of an entity, else it returns 0.0.",
	},
	{
		Name:        "ride_head_y_rotation",
		Signature:   "(maxRotation?: number): number",
		Description: "Takes one optional argument as a parameter. Returns the head y world-rotation of the ride of an entity, else it returns 0.0. First parameter only for horses, zombie horses, skeleton horses, donkeys and mules that clamps rotation in degrees.",
	},
	{
		Name:        "rider_body_x_rotation",
		Signature:   "(index: number): number",
		Description: "Returns the body pitch world-rotation of a valid rider at the provided index if called on an entity, else it returns 0.0.",
	},
	{
		Name:        "rider_body_y_rotation",
		Signature:   "(index: number): number",
		Description: "Returns the body yaw world-rotation of a valid rider at the provided index if called on an entity, else it returns 0.0.",
	},
	{
		Name:        "rider_head_x_rotation",
		Signature:   "(index: number): number",
		Description: "Takes one argument as a parameter. Returns the head x world-rotation of the rider entity at the provided index, else it returns 0.0.",
	},
	{
		Name:        "rider_head_y_rotation",
		Signature:   "(index: number, maxRotation?: number): number",
		Description: "Takes one or two arguments as parameters. Returns the head y world-rotation of the rider entity at the provided index, else it returns 0.0. Horses, zombie horses, skeleton horses, donkeys and mules require a second parameter that clamps rotation in degrees.",
	},
	{
		Name:        "roll_counter",
		Signature:   ": number",
		Description: "Returns the roll counter of the entity.",
	},
	{
		Name:        "rotation_to_camera",
		Signature:   "(axis: number): number",
		Description: "Returns the rotation required to aim at the camera. Requires one argument representing the rotation axis you would like (0 for x, 1 for y).",
	},
	{
		Name:        "scoreboard",
		Signature:   "(objective: string): number",
		Description: "Takes one argument - the name of the scoreboard entry for this entity. Returns the specified scoreboard value for this entity. Available only with behavior packs.",
	},
	{
		Name:        "server_memory_tier",
		Signature:   ": number",
		Description: "Returns a number representing the server RAM memory tier, 0 = 'SuperLow', 1 = 'Low', 2 = 'Mid', 3 = 'High', or 4 = 'SuperHigh'. Available on the server side (Behavior Packs) only.",
	},
	{
		Name:        "shake_angle",
		Signature:   ": number",
		Description: "Returns the shaking angle of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "shake_time",
		Signature:   ": number",
		Description: "Returns the shake time of the entity.",
	},
	{
		Name:        "shield_blocking_bob",
		Signature:   ": number",
		Description: "Returns the how much the offhand shield should translate down when blocking and being hit.",
	},
	{
		Name:        "show_bottom",
		Signature:   ": boolean",
		Description: "Returns 1.0 if we render the entity's bottom, else it returns 0.0.",
	},
	{
		Name:        "sit_amount",
		Signature:   ": number",
		Description: "Returns the current sit amount of the entity.",
	},
	{
		Name:        "skin_id",
		Signature:   ": number",
		Description: "Returns the entity's skin ID",
	},
	{
		Name:        "sleep_rotation",
		Signature:   ": number",
		Description: "Returns the rotation of the bed the player is sleeping on.",
	},
	{
		Name:        "sneeze_counter",
		Signature:   ": number",
		Description: "Returns the sneeze counter of the entity.",
	},
	{
		Name:        "spellcolor",
		Signature:   ": rgba",
		Description: "Returns a struct representing the entity spell color for the specified entity. The struct contains '.r' '.g' '.b' and '.a' members, each 0.0 to 1.0. If no actor is specified, each member value will be 0.0.",
	},
	{
		Name:        "standing_scale",
		Signature:   ": number",
		Description: "Returns the scale of how standing up the entity is.",
	},
	{
		Name:        "state_time",
		Signature:   ": number",
		Description: "Only valid in an animation controller. Returns the time in seconds in the current animation controller state.",
	},
	{
		Name:        "structural_integrity",
		Signature:   ": number",
		Description: "Returns the structural integrity for the actor, otherwise returns 0.",
	},
	{
		Name:        "surface_particle_color",
		Signature:   ": rgba",
		Description: "Returns the particle color for the block located in the surface below the actor (scanned up to 10 blocks down). The struct contains '.r' '.g' '.b' and '.a' members, each 0.0 to 1.0. If no actor is specified or if no surface is found, each member value is set to 0.0. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "surface_particle_texture_coordinate",
		Signature:   ": { u: number, v: number }",
		Description: "Returns the texture coordinate for generating particles for the block located in the surface below the actor (scanned up to 10 blocks down) in a struct with 'u' and 'v' keys. If no actor is specified or if no surface is found, u and v will be 0.0. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "surface_particle_texture_size",
		Signature:   ": { u: number, v: number }",
		Description: "Returns the texture size for generating particles for the block located in the surface below the actor (scanned up to 10 blocks down). If no actor is specified or if no surface is found, each member value will be 0.0. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "swell_amount",
		Signature:   ": number",
		Description: "Returns how swollen the entity is.",
	},
	{
		Name:        "swelling_dir",
		Signature:   ": number",
		Description: "Returns the swelling direction of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "swim_amount",
		Signature:   ": number",
		Description: "Returns the amount the current entity is swimming.",
	},
	{
		Name:        "tail_angle",
		Signature:   ": number",
		Description: "Returns the angle of the tail of the entity if it makes sense, else it returns 0.0.",
	},
	{
		Name:        "target_x_rotation",
		Signature:   ": number",
		Description: "Returns the x rotation required to aim at the entity's current target if it has one, else it returns 0.0.",
	},
	{
		Name:        "target_y_rotation",
		Signature:   ": number",
		Description: "Returns the y rotation required to aim at the entity's current target if it has one, else it returns 0.0.",
	},
	{
		Name:        "texture_frame_index",
		Signature:   ": number",
		Description: "Returns the icon index of the experience orb.",
	},
	{
		Name:        "time_of_day",
		Signature:   ": number",
		Description: "Returns the time of day (midnight=0.0, sunrise=0.25, noon=0.5, sunset=0.75) of the dimension the entity is in.",
	},
	{
		Name:        "time_since_last_vibration_detection",
		Signature:   ": number",
		Description: "Returns the time in seconds since the last vibration detected by the actor. On errors or if no vibration has been detected yet, returns -1. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "time_stamp",
		Signature:   ": number",
		Description: "Returns the current time stamp of the level",
	},
	{
		Name:        "timer_flag_1",
		Signature:   ": number",
		Description: "Returns 1.0 if behavior.timer_flag_1 is running, else it returns 0.0.",
	},
	{
		Name:        "timer_flag_2",
		Signature:   ": number",
		Description: "Returns 1.0 if behavior.timer_flag_2 is running, else it returns 0.0.",
	},
	{
		Name:        "timer_flag_3",
		Signature:   ": number",
		Description: "Returns 1.0 if behavior.timer_flag_3 is running, else it returns 0.0.",
	},
	{
		Name:        "total_emitter_count",
		Signature:   ": number",
		Description: "Returns the total number of active emitters in the world.",
	},
	{
		Name:        "total_particle_count",
		Signature:   ": number",
		Description: "Returns the total number of active particles in the world.",
	},
	{
		Name:        "touch_only_affects_hotbar",
		Signature:   ": boolean",
		Description: "Returns 1.0 if the touch input only affects the touchbar, otherwise returns 0.0. Available on the Client (Resource Packs) only.",
	},
	{
		Name:        "trade_tier",
		Signature:   ": number",
		Description: "Returns the trade tier of the entity if it makes sense, else it returns 0.0",
	},
	{
		Name:        "unhappy_counter",
		Signature:   ": number",
		Description: "Always returns zero. (Was originally meant to indicate Panda unhappiness but due to an early code change it has always only returned zero)",
	},
	{
		Name:        "variant",
		Signature:   ": number",
		Description: "Returns the entity's variant index",
	},
	{
		Name:        "vertical_speed",
		Signature:   ": number",
		Description: "Returns the speed of the entity up or down in meters/second, where positive is up.",
	},
	{
		Name:        "walk_distance",
		Signature:   ": number",
		Description: "Returns the total distance traveled by an entity while on the ground and not sneaking.",
	},
	{
		Name:        "wing_flap_position",
		Signature:   ": number",
		Description: "Returns the wing flap position of the entity, or 0.0 if this doesn't make sense.",
	},
	{
		Name:        "wing_flap_speed",
		Signature:   ": number",
		Description: "Returns the wing flap speed of the entity, or 0.0 if this doesn't make sense.",
	},
	{
		Name:        "yaw_speed",
		Signature:   ": number",
		Description: "Returns the entity's yaw speed",
	},
}

var molangMath = []Method{
	{
		Name:        "abs",
		Signature:   "(value: number): number",
		Description: "Absolute value of value",
	},
	{
		Name:        "acos",
		Signature:   "(value: number): number",
		Description: "arccos of value",
	},
	{
		Name:        "asin",
		Signature:   "(value: number): number",
		Description: "arcsin of value",
	},
	{
		Name:        "atan",
		Signature:   "(value: number): number",
		Description: "arctan of value",
	},
	{
		Name:        "atan2",
		Signature:   "(y: number, x: number): number",
		Description: "arctan of y/x. NOTE: the order of arguments!",
	},
	{
		Name:        "ceil",
		Signature:   "(value: number): number",
		Description: "Round value up to nearest integral number",
	},
	{
		Name:        "clamp",
		Signature:   "(value: number, min: number, max: number): number",
		Description: "Clamp value to between min and max inclusive",
	},
	{
		Name:        "cos",
		Signature:   "(value: number): number",
		Description: "Cosine (in degrees) of value",
	},
	{
		Name:        "die_roll",
		Signature:   "(num: number, low: number, high: number): number",
		Description: "returns the sum of 'num' random numbers, each with a value from low to high`. Note: the generated random numbers are not integers like normal dice. For that, use `math.die_roll_integer`.",
	},
	{
		Name:        "die_roll_integer",
		Signature:   "(num: number, low: number, high: number): number",
		Description: "returns the sum of 'num' random integer numbers, each with a value from low to high`. Note: the generated random numbers are integers like normal dice.",
	},
	{
		Name:        "exp",
		Signature:   "(value: number): number",
		Description: "Calculates e to the value'th power",
	},
	{
		Name:        "floor",
		Signature:   "(value: number): number",
		Description: "Round value down to nearest integral number",
	},
	{
		Name:        "hermite_blend",
		Signature:   "(value: number): number",
		Description: "Useful for simple smooth curve interpolation using one of the Hermite Basis functions: `3t^2 - 2t^3`. Note that while any valid float is a valid input, this function works best in the range [0,1].",
	},
	{
		Name:        "lerp",
		Signature:   "(start: number, end: number, 0_to_1: number): number",
		Description: "Lerp from start to end via 0_to_1",
	},
	{
		Name:        "lerprotate",
		Signature:   "(start: number, end: number, 0_to_1: number): number",
		Description: "Lerp the shortest direction around a circle from start degrees to end degrees via 0_to_1",
	},
	{
		Name:        "ln",
		Signature:   "(value: number): number",
		Description: "Natural logarithm of value",
	},
	{
		Name:        "max",
		Signature:   "(A: number, B: number): number",
		Description: "Return highest value of A or B",
	},
	{
		Name:        "min",
		Signature:   "(A: number, B: number): number",
		Description: "Return lowest value of A or B",
	},
	{
		Name:        "min_angle",
		Signature:   "(value: number): number",
		Description: "Minimize angle magnitude (in degrees) into the range [-180, 180]",
	},
	{
		Name:        "mod",
		Signature:   "(value: number, denominator: number): number",
		Description: "Return the remainder of value / denominator",
	},
	{
		Name:        "pi",
		Signature:   ": number",
		Description: "Returns the float representation of the constant pi.",
	},
	{
		Name:        "pow",
		Signature:   "(base: number, exponent: number): number",
		Description: "Elevates `base` to the `exponent`'th power",
	},
	{
		Name:        "random",
		Signature:   "(low: number, high: number): number",
		Description: "Random value between low and high inclusive",
	},
	{
		Name:        "random_integer",
		Signature:   "(low: number, high: number): number",
		Description: "Random integer value between low and high inclusive",
	},
	{
		Name:        "round",
		Signature:   "(value: number): number",
		Description: "Round value to nearest integral number",
	},
	{
		Name:        "sin",
		Signature:   "(value: number): number",
		Description: "Sine (in degrees) of value",
	},
	{
		Name:        "sqrt",
		Signature:   "(value: number): number",
		Description: "Square root of value",
	},
	{
		Name:        "trunc",
		Signature:   "(value: number): number",
		Description: "Round value towards zero",
	},
}

func GetMethodList(prefix string) []Method {
	if match, _ := regexp.MatchString(`(?i)(q|query)`, prefix); match {
		return molangQueries
	}
	if match, _ := regexp.MatchString(`(?i)math`, prefix); match {
		return molangMath
	}
	return nil
}

func GetMethod(prefix string, name string) (Method, bool) {
	name = strings.TrimPrefix(name, ".")
	return sliceutil.Find(GetMethodList(prefix), func(method Method) bool {
		return method.Name == name
	})
}
