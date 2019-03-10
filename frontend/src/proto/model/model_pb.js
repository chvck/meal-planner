/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.chvck.mealplanner.model.Ingredient', null, global);
goog.exportSymbol('proto.chvck.mealplanner.model.Planner', null, global);
goog.exportSymbol('proto.chvck.mealplanner.model.Planner.Mealtime', null, global);
goog.exportSymbol('proto.chvck.mealplanner.model.Recipe', null, global);
goog.exportSymbol('proto.chvck.mealplanner.model.User', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.chvck.mealplanner.model.User = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.chvck.mealplanner.model.User, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.chvck.mealplanner.model.User.displayName = 'proto.chvck.mealplanner.model.User';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.chvck.mealplanner.model.Ingredient = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.chvck.mealplanner.model.Ingredient, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.chvck.mealplanner.model.Ingredient.displayName = 'proto.chvck.mealplanner.model.Ingredient';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.chvck.mealplanner.model.Recipe = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.chvck.mealplanner.model.Recipe.repeatedFields_, null);
};
goog.inherits(proto.chvck.mealplanner.model.Recipe, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.chvck.mealplanner.model.Recipe.displayName = 'proto.chvck.mealplanner.model.Recipe';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.chvck.mealplanner.model.Planner = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.chvck.mealplanner.model.Planner.repeatedFields_, null);
};
goog.inherits(proto.chvck.mealplanner.model.Planner, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.chvck.mealplanner.model.Planner.displayName = 'proto.chvck.mealplanner.model.Planner';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.chvck.mealplanner.model.User.prototype.toObject = function(opt_includeInstance) {
  return proto.chvck.mealplanner.model.User.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.chvck.mealplanner.model.User} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.User.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = jspb.Message.getFieldWithDefault(msg, 1, "")) == null ? undefined : f,
    username: (f = jspb.Message.getFieldWithDefault(msg, 2, "")) == null ? undefined : f,
    email: (f = jspb.Message.getFieldWithDefault(msg, 3, "")) == null ? undefined : f,
    createdAt: (f = jspb.Message.getFieldWithDefault(msg, 4, 0)) == null ? undefined : f,
    updatedAt: (f = jspb.Message.getFieldWithDefault(msg, 5, 0)) == null ? undefined : f,
    lastLogin: (f = jspb.Message.getFieldWithDefault(msg, 6, 0)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.chvck.mealplanner.model.User}
 */
proto.chvck.mealplanner.model.User.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.chvck.mealplanner.model.User;
  return proto.chvck.mealplanner.model.User.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.chvck.mealplanner.model.User} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.chvck.mealplanner.model.User}
 */
proto.chvck.mealplanner.model.User.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreatedAt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUpdatedAt(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLastLogin(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.chvck.mealplanner.model.User.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.chvck.mealplanner.model.User.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.chvck.mealplanner.model.User} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.User.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCreatedAt();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getUpdatedAt();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getLastLogin();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.chvck.mealplanner.model.User.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.User.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string username = 2;
 * @return {string}
 */
proto.chvck.mealplanner.model.User.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.User.prototype.setUsername = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string email = 3;
 * @return {string}
 */
proto.chvck.mealplanner.model.User.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.User.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int64 created_at = 4;
 * @return {number}
 */
proto.chvck.mealplanner.model.User.prototype.getCreatedAt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.User.prototype.setCreatedAt = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 updated_at = 5;
 * @return {number}
 */
proto.chvck.mealplanner.model.User.prototype.getUpdatedAt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.User.prototype.setUpdatedAt = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 last_login = 6;
 * @return {number}
 */
proto.chvck.mealplanner.model.User.prototype.getLastLogin = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.User.prototype.setLastLogin = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.toObject = function(opt_includeInstance) {
  return proto.chvck.mealplanner.model.Ingredient.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.chvck.mealplanner.model.Ingredient} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Ingredient.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = jspb.Message.getFieldWithDefault(msg, 1, "")) == null ? undefined : f,
    recipeId: (f = jspb.Message.getFieldWithDefault(msg, 2, "")) == null ? undefined : f,
    name: (f = jspb.Message.getFieldWithDefault(msg, 3, "")) == null ? undefined : f,
    measure: (f = jspb.Message.getFieldWithDefault(msg, 4, "")) == null ? undefined : f,
    quantity: (f = jspb.Message.getFieldWithDefault(msg, 5, "")) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.chvck.mealplanner.model.Ingredient}
 */
proto.chvck.mealplanner.model.Ingredient.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.chvck.mealplanner.model.Ingredient;
  return proto.chvck.mealplanner.model.Ingredient.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.chvck.mealplanner.model.Ingredient} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.chvck.mealplanner.model.Ingredient}
 */
proto.chvck.mealplanner.model.Ingredient.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRecipeId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasure(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setQuantity(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.chvck.mealplanner.model.Ingredient.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.chvck.mealplanner.model.Ingredient} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Ingredient.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRecipeId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getMeasure();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getQuantity();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Ingredient.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string recipe_id = 2;
 * @return {string}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.getRecipeId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Ingredient.prototype.setRecipeId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Ingredient.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string measure = 4;
 * @return {string}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.getMeasure = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Ingredient.prototype.setMeasure = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string quantity = 5;
 * @return {string}
 */
proto.chvck.mealplanner.model.Ingredient.prototype.getQuantity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Ingredient.prototype.setQuantity = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.chvck.mealplanner.model.Recipe.repeatedFields_ = [9];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.chvck.mealplanner.model.Recipe.prototype.toObject = function(opt_includeInstance) {
  return proto.chvck.mealplanner.model.Recipe.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.chvck.mealplanner.model.Recipe} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Recipe.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = jspb.Message.getFieldWithDefault(msg, 1, "")) == null ? undefined : f,
    userId: (f = jspb.Message.getFieldWithDefault(msg, 2, "")) == null ? undefined : f,
    name: (f = jspb.Message.getFieldWithDefault(msg, 3, "")) == null ? undefined : f,
    instructions: (f = jspb.Message.getFieldWithDefault(msg, 4, "")) == null ? undefined : f,
    yield: (f = jspb.Message.getFieldWithDefault(msg, 5, 0)) == null ? undefined : f,
    prepTime: (f = jspb.Message.getFieldWithDefault(msg, 6, 0)) == null ? undefined : f,
    cookTime: (f = jspb.Message.getFieldWithDefault(msg, 7, 0)) == null ? undefined : f,
    description: (f = jspb.Message.getFieldWithDefault(msg, 8, "")) == null ? undefined : f,
    ingredientsList: jspb.Message.toObjectList(msg.getIngredientsList(),
    proto.chvck.mealplanner.model.Ingredient.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.chvck.mealplanner.model.Recipe}
 */
proto.chvck.mealplanner.model.Recipe.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.chvck.mealplanner.model.Recipe;
  return proto.chvck.mealplanner.model.Recipe.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.chvck.mealplanner.model.Recipe} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.chvck.mealplanner.model.Recipe}
 */
proto.chvck.mealplanner.model.Recipe.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setInstructions(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setYield(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPrepTime(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCookTime(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 9:
      var value = new proto.chvck.mealplanner.model.Ingredient;
      reader.readMessage(value,proto.chvck.mealplanner.model.Ingredient.deserializeBinaryFromReader);
      msg.addIngredients(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.chvck.mealplanner.model.Recipe.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.chvck.mealplanner.model.Recipe.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.chvck.mealplanner.model.Recipe} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Recipe.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getInstructions();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getYield();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getPrepTime();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getCookTime();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getIngredientsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      9,
      f,
      proto.chvck.mealplanner.model.Ingredient.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Recipe.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Recipe.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Recipe.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string instructions = 4;
 * @return {string}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getInstructions = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Recipe.prototype.setInstructions = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 yield = 5;
 * @return {number}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getYield = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.Recipe.prototype.setYield = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 prep_time = 6;
 * @return {number}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getPrepTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.Recipe.prototype.setPrepTime = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 cook_time = 7;
 * @return {number}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getCookTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.Recipe.prototype.setCookTime = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional string description = 8;
 * @return {string}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Recipe.prototype.setDescription = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * repeated Ingredient ingredients = 9;
 * @return {!Array<!proto.chvck.mealplanner.model.Ingredient>}
 */
proto.chvck.mealplanner.model.Recipe.prototype.getIngredientsList = function() {
  return /** @type{!Array<!proto.chvck.mealplanner.model.Ingredient>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.chvck.mealplanner.model.Ingredient, 9));
};


/** @param {!Array<!proto.chvck.mealplanner.model.Ingredient>} value */
proto.chvck.mealplanner.model.Recipe.prototype.setIngredientsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 9, value);
};


/**
 * @param {!proto.chvck.mealplanner.model.Ingredient=} opt_value
 * @param {number=} opt_index
 * @return {!proto.chvck.mealplanner.model.Ingredient}
 */
proto.chvck.mealplanner.model.Recipe.prototype.addIngredients = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 9, opt_value, proto.chvck.mealplanner.model.Ingredient, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.chvck.mealplanner.model.Recipe.prototype.clearIngredientsList = function() {
  this.setIngredientsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.chvck.mealplanner.model.Planner.repeatedFields_ = [5];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.chvck.mealplanner.model.Planner.prototype.toObject = function(opt_includeInstance) {
  return proto.chvck.mealplanner.model.Planner.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.chvck.mealplanner.model.Planner} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Planner.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = jspb.Message.getFieldWithDefault(msg, 1, "")) == null ? undefined : f,
    userId: (f = jspb.Message.getFieldWithDefault(msg, 2, "")) == null ? undefined : f,
    date: (f = jspb.Message.getFieldWithDefault(msg, 3, 0)) == null ? undefined : f,
    mealtime: (f = jspb.Message.getFieldWithDefault(msg, 4, 0)) == null ? undefined : f,
    recipeIdsList: (f = jspb.Message.getRepeatedField(msg, 5)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.chvck.mealplanner.model.Planner}
 */
proto.chvck.mealplanner.model.Planner.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.chvck.mealplanner.model.Planner;
  return proto.chvck.mealplanner.model.Planner.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.chvck.mealplanner.model.Planner} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.chvck.mealplanner.model.Planner}
 */
proto.chvck.mealplanner.model.Planner.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDate(value);
      break;
    case 4:
      var value = /** @type {!proto.chvck.mealplanner.model.Planner.Mealtime} */ (reader.readEnum());
      msg.setMealtime(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.addRecipeIds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.chvck.mealplanner.model.Planner.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.chvck.mealplanner.model.Planner.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.chvck.mealplanner.model.Planner} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.chvck.mealplanner.model.Planner.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDate();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getMealtime();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getRecipeIdsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      5,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.chvck.mealplanner.model.Planner.Mealtime = {
  BREAKFAST: 0,
  LUNCH: 1,
  TEA: 2,
  SUPPER: 3,
  SNACK: 4
};

/**
 * optional string id = 1;
 * @return {string}
 */
proto.chvck.mealplanner.model.Planner.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Planner.prototype.setId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.chvck.mealplanner.model.Planner.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.chvck.mealplanner.model.Planner.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 date = 3;
 * @return {number}
 */
proto.chvck.mealplanner.model.Planner.prototype.getDate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.chvck.mealplanner.model.Planner.prototype.setDate = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional Mealtime mealtime = 4;
 * @return {!proto.chvck.mealplanner.model.Planner.Mealtime}
 */
proto.chvck.mealplanner.model.Planner.prototype.getMealtime = function() {
  return /** @type {!proto.chvck.mealplanner.model.Planner.Mealtime} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {!proto.chvck.mealplanner.model.Planner.Mealtime} value */
proto.chvck.mealplanner.model.Planner.prototype.setMealtime = function(value) {
  jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * repeated string recipe_ids = 5;
 * @return {!Array<string>}
 */
proto.chvck.mealplanner.model.Planner.prototype.getRecipeIdsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 5));
};


/** @param {!Array<string>} value */
proto.chvck.mealplanner.model.Planner.prototype.setRecipeIdsList = function(value) {
  jspb.Message.setField(this, 5, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 */
proto.chvck.mealplanner.model.Planner.prototype.addRecipeIds = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 5, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.chvck.mealplanner.model.Planner.prototype.clearRecipeIdsList = function() {
  this.setRecipeIdsList([]);
};


goog.object.extend(exports, proto.chvck.mealplanner.model);
