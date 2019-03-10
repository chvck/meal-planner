// package: chvck.mealplanner.service
// file: proto/service/service.proto

import * as jspb from "google-protobuf";
import * as proto_model_model_pb from "../../proto/model/model_pb";

export class AllRecipesRequest extends jspb.Message {
  getOffset(): number;
  setOffset(value: number): void;

  getLimit(): number;
  setLimit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AllRecipesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AllRecipesRequest): AllRecipesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AllRecipesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AllRecipesRequest;
  static deserializeBinaryFromReader(message: AllRecipesRequest, reader: jspb.BinaryReader): AllRecipesRequest;
}

export namespace AllRecipesRequest {
  export type AsObject = {
    offset: number,
    limit: number,
  }
}

export class AllRecipesResponse extends jspb.Message {
  clearRecipesList(): void;
  getRecipesList(): Array<proto_model_model_pb.Recipe>;
  setRecipesList(value: Array<proto_model_model_pb.Recipe>): void;
  addRecipes(value?: proto_model_model_pb.Recipe, index?: number): proto_model_model_pb.Recipe;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AllRecipesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AllRecipesResponse): AllRecipesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AllRecipesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AllRecipesResponse;
  static deserializeBinaryFromReader(message: AllRecipesResponse, reader: jspb.BinaryReader): AllRecipesResponse;
}

export namespace AllRecipesResponse {
  export type AsObject = {
    recipesList: Array<proto_model_model_pb.Recipe.AsObject>,
  }
}

export class RecipeByIDRequest extends jspb.Message {
  getRecipeId(): number;
  setRecipeId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecipeByIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RecipeByIDRequest): RecipeByIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RecipeByIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecipeByIDRequest;
  static deserializeBinaryFromReader(message: RecipeByIDRequest, reader: jspb.BinaryReader): RecipeByIDRequest;
}

export namespace RecipeByIDRequest {
  export type AsObject = {
    recipeId: number,
  }
}

export class RecipeByIDResponse extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): proto_model_model_pb.Recipe | undefined;
  setRecipe(value?: proto_model_model_pb.Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RecipeByIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RecipeByIDResponse): RecipeByIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RecipeByIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RecipeByIDResponse;
  static deserializeBinaryFromReader(message: RecipeByIDResponse, reader: jspb.BinaryReader): RecipeByIDResponse;
}

export namespace RecipeByIDResponse {
  export type AsObject = {
    recipe?: proto_model_model_pb.Recipe.AsObject,
  }
}

export class CreateRecipeRequest extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): proto_model_model_pb.Recipe | undefined;
  setRecipe(value?: proto_model_model_pb.Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRecipeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRecipeRequest): CreateRecipeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRecipeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRecipeRequest;
  static deserializeBinaryFromReader(message: CreateRecipeRequest, reader: jspb.BinaryReader): CreateRecipeRequest;
}

export namespace CreateRecipeRequest {
  export type AsObject = {
    recipe?: proto_model_model_pb.Recipe.AsObject,
  }
}

export class CreateRecipeResponse extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): proto_model_model_pb.Recipe | undefined;
  setRecipe(value?: proto_model_model_pb.Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRecipeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRecipeResponse): CreateRecipeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRecipeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRecipeResponse;
  static deserializeBinaryFromReader(message: CreateRecipeResponse, reader: jspb.BinaryReader): CreateRecipeResponse;
}

export namespace CreateRecipeResponse {
  export type AsObject = {
    recipe?: proto_model_model_pb.Recipe.AsObject,
  }
}

export class UpdateRecipeRequest extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): proto_model_model_pb.Recipe | undefined;
  setRecipe(value?: proto_model_model_pb.Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRecipeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRecipeRequest): UpdateRecipeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRecipeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRecipeRequest;
  static deserializeBinaryFromReader(message: UpdateRecipeRequest, reader: jspb.BinaryReader): UpdateRecipeRequest;
}

export namespace UpdateRecipeRequest {
  export type AsObject = {
    recipe?: proto_model_model_pb.Recipe.AsObject,
  }
}

export class UpdateRecipeResponse extends jspb.Message {
  hasRecipe(): boolean;
  clearRecipe(): void;
  getRecipe(): proto_model_model_pb.Recipe | undefined;
  setRecipe(value?: proto_model_model_pb.Recipe): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRecipeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRecipeResponse): UpdateRecipeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRecipeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRecipeResponse;
  static deserializeBinaryFromReader(message: UpdateRecipeResponse, reader: jspb.BinaryReader): UpdateRecipeResponse;
}

export namespace UpdateRecipeResponse {
  export type AsObject = {
    recipe?: proto_model_model_pb.Recipe.AsObject,
  }
}

export class DeleteRecipeRequest extends jspb.Message {
  getRecipeId(): number;
  setRecipeId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRecipeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRecipeRequest): DeleteRecipeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRecipeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRecipeRequest;
  static deserializeBinaryFromReader(message: DeleteRecipeRequest, reader: jspb.BinaryReader): DeleteRecipeRequest;
}

export namespace DeleteRecipeRequest {
  export type AsObject = {
    recipeId: number,
  }
}

export class DeleteRecipeResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRecipeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRecipeResponse): DeleteRecipeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRecipeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRecipeResponse;
  static deserializeBinaryFromReader(message: DeleteRecipeResponse, reader: jspb.BinaryReader): DeleteRecipeResponse;
}

export namespace DeleteRecipeResponse {
  export type AsObject = {
  }
}

export class LoginUserRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginUserRequest): LoginUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginUserRequest;
  static deserializeBinaryFromReader(message: LoginUserRequest, reader: jspb.BinaryReader): LoginUserRequest;
}

export namespace LoginUserRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginUserResponse extends jspb.Message {
  getToken(): Uint8Array | string;
  getToken_asU8(): Uint8Array;
  getToken_asB64(): string;
  setToken(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginUserResponse): LoginUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginUserResponse;
  static deserializeBinaryFromReader(message: LoginUserResponse, reader: jspb.BinaryReader): LoginUserResponse;
}

export namespace LoginUserResponse {
  export type AsObject = {
    token: Uint8Array | string,
  }
}

export class CreateUserRequest extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): proto_model_model_pb.User | undefined;
  setUser(value?: proto_model_model_pb.User): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    user?: proto_model_model_pb.User.AsObject,
    password: string,
  }
}

export class CreateUserResponse extends jspb.Message {
  hasUser(): boolean;
  clearUser(): void;
  getUser(): proto_model_model_pb.User | undefined;
  setUser(value?: proto_model_model_pb.User): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
    user?: proto_model_model_pb.User.AsObject,
  }
}

