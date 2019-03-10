// package: chvck.mealplanner.service
// file: proto/service/service.proto

import * as proto_service_service_pb from "../../proto/service/service_pb";
import {grpc} from "@improbable-eng/grpc-web";

type MealPlannerServiceAllRecipes = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.AllRecipesRequest;
  readonly responseType: typeof proto_service_service_pb.AllRecipesResponse;
};

type MealPlannerServiceRecipeByID = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.RecipeByIDRequest;
  readonly responseType: typeof proto_service_service_pb.RecipeByIDResponse;
};

type MealPlannerServiceCreateRecipe = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.CreateRecipeRequest;
  readonly responseType: typeof proto_service_service_pb.CreateRecipeResponse;
};

type MealPlannerServiceUpdateRecipe = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.UpdateRecipeRequest;
  readonly responseType: typeof proto_service_service_pb.UpdateRecipeResponse;
};

type MealPlannerServiceDeleteRecipe = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.DeleteRecipeRequest;
  readonly responseType: typeof proto_service_service_pb.DeleteRecipeResponse;
};

type MealPlannerServiceCreateUser = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.CreateUserRequest;
  readonly responseType: typeof proto_service_service_pb.CreateUserResponse;
};

type MealPlannerServiceLoginUser = {
  readonly methodName: string;
  readonly service: typeof MealPlannerService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_service_service_pb.LoginUserRequest;
  readonly responseType: typeof proto_service_service_pb.LoginUserResponse;
};

export class MealPlannerService {
  static readonly serviceName: string;
  static readonly AllRecipes: MealPlannerServiceAllRecipes;
  static readonly RecipeByID: MealPlannerServiceRecipeByID;
  static readonly CreateRecipe: MealPlannerServiceCreateRecipe;
  static readonly UpdateRecipe: MealPlannerServiceUpdateRecipe;
  static readonly DeleteRecipe: MealPlannerServiceDeleteRecipe;
  static readonly CreateUser: MealPlannerServiceCreateUser;
  static readonly LoginUser: MealPlannerServiceLoginUser;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class MealPlannerServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  allRecipes(
    requestMessage: proto_service_service_pb.AllRecipesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.AllRecipesResponse|null) => void
  ): UnaryResponse;
  allRecipes(
    requestMessage: proto_service_service_pb.AllRecipesRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.AllRecipesResponse|null) => void
  ): UnaryResponse;
  recipeByID(
    requestMessage: proto_service_service_pb.RecipeByIDRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.RecipeByIDResponse|null) => void
  ): UnaryResponse;
  recipeByID(
    requestMessage: proto_service_service_pb.RecipeByIDRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.RecipeByIDResponse|null) => void
  ): UnaryResponse;
  createRecipe(
    requestMessage: proto_service_service_pb.CreateRecipeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.CreateRecipeResponse|null) => void
  ): UnaryResponse;
  createRecipe(
    requestMessage: proto_service_service_pb.CreateRecipeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.CreateRecipeResponse|null) => void
  ): UnaryResponse;
  updateRecipe(
    requestMessage: proto_service_service_pb.UpdateRecipeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.UpdateRecipeResponse|null) => void
  ): UnaryResponse;
  updateRecipe(
    requestMessage: proto_service_service_pb.UpdateRecipeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.UpdateRecipeResponse|null) => void
  ): UnaryResponse;
  deleteRecipe(
    requestMessage: proto_service_service_pb.DeleteRecipeRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.DeleteRecipeResponse|null) => void
  ): UnaryResponse;
  deleteRecipe(
    requestMessage: proto_service_service_pb.DeleteRecipeRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.DeleteRecipeResponse|null) => void
  ): UnaryResponse;
  createUser(
    requestMessage: proto_service_service_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.CreateUserResponse|null) => void
  ): UnaryResponse;
  createUser(
    requestMessage: proto_service_service_pb.CreateUserRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.CreateUserResponse|null) => void
  ): UnaryResponse;
  loginUser(
    requestMessage: proto_service_service_pb.LoginUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.LoginUserResponse|null) => void
  ): UnaryResponse;
  loginUser(
    requestMessage: proto_service_service_pb.LoginUserRequest,
    callback: (error: ServiceError|null, responseMessage: proto_service_service_pb.LoginUserResponse|null) => void
  ): UnaryResponse;
}

