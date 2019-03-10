/* eslint-disable */
// package: chvck.mealplanner.service
// file: proto/service/service.proto

var proto_service_service_pb = require("../../proto/service/service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var MealPlannerService = (function () {
  function MealPlannerService() {}
  MealPlannerService.serviceName = "chvck.mealplanner.service.MealPlannerService";
  return MealPlannerService;
}());

MealPlannerService.AllRecipes = {
  methodName: "AllRecipes",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.AllRecipesRequest,
  responseType: proto_service_service_pb.AllRecipesResponse
};

MealPlannerService.RecipeByID = {
  methodName: "RecipeByID",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.RecipeByIDRequest,
  responseType: proto_service_service_pb.RecipeByIDResponse
};

MealPlannerService.CreateRecipe = {
  methodName: "CreateRecipe",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.CreateRecipeRequest,
  responseType: proto_service_service_pb.CreateRecipeResponse
};

MealPlannerService.UpdateRecipe = {
  methodName: "UpdateRecipe",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.UpdateRecipeRequest,
  responseType: proto_service_service_pb.UpdateRecipeResponse
};

MealPlannerService.DeleteRecipe = {
  methodName: "DeleteRecipe",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.DeleteRecipeRequest,
  responseType: proto_service_service_pb.DeleteRecipeResponse
};

MealPlannerService.CreateUser = {
  methodName: "CreateUser",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.CreateUserRequest,
  responseType: proto_service_service_pb.CreateUserResponse
};

MealPlannerService.LoginUser = {
  methodName: "LoginUser",
  service: MealPlannerService,
  requestStream: false,
  responseStream: false,
  requestType: proto_service_service_pb.LoginUserRequest,
  responseType: proto_service_service_pb.LoginUserResponse
};

exports.MealPlannerService = MealPlannerService;

function MealPlannerServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MealPlannerServiceClient.prototype.allRecipes = function allRecipes(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.AllRecipes, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.recipeByID = function recipeByID(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.RecipeByID, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.createRecipe = function createRecipe(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.CreateRecipe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.updateRecipe = function updateRecipe(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.UpdateRecipe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.deleteRecipe = function deleteRecipe(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.DeleteRecipe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.createUser = function createUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.CreateUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MealPlannerServiceClient.prototype.loginUser = function loginUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MealPlannerService.LoginUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.MealPlannerServiceClient = MealPlannerServiceClient;

