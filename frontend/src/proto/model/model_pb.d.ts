// package: chvck.mealplanner.model
// file: proto/model/model.proto

import * as jspb from "google-protobuf";

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getCreatedAt(): number;
  setCreatedAt(value: number): void;

  getUpdatedAt(): number;
  setUpdatedAt(value: number): void;

  getLastLogin(): number;
  setLastLogin(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    username: string,
    email: string,
    createdAt: number,
    updatedAt: number,
    lastLogin: number,
  }
}

export class Ingredient extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getRecipeId(): string;
  setRecipeId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getMeasure(): string;
  setMeasure(value: string): void;

  getQuantity(): string;
  setQuantity(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ingredient.AsObject;
  static toObject(includeInstance: boolean, msg: Ingredient): Ingredient.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Ingredient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ingredient;
  static deserializeBinaryFromReader(message: Ingredient, reader: jspb.BinaryReader): Ingredient;
}

export namespace Ingredient {
  export type AsObject = {
    id: string,
    recipeId: string,
    name: string,
    measure: string,
    quantity: string,
  }
}

export class Recipe extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUserId(): string;
  setUserId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getInstructions(): string;
  setInstructions(value: string): void;

  getYield(): number;
  setYield(value: number): void;

  getPrepTime(): number;
  setPrepTime(value: number): void;

  getCookTime(): number;
  setCookTime(value: number): void;

  getDescription(): string;
  setDescription(value: string): void;

  clearIngredientsList(): void;
  getIngredientsList(): Array<Ingredient>;
  setIngredientsList(value: Array<Ingredient>): void;
  addIngredients(value?: Ingredient, index?: number): Ingredient;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Recipe.AsObject;
  static toObject(includeInstance: boolean, msg: Recipe): Recipe.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Recipe, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Recipe;
  static deserializeBinaryFromReader(message: Recipe, reader: jspb.BinaryReader): Recipe;
}

export namespace Recipe {
  export type AsObject = {
    id: string,
    userId: string,
    name: string,
    instructions: string,
    yield: number,
    prepTime: number,
    cookTime: number,
    description: string,
    ingredientsList: Array<Ingredient.AsObject>,
  }
}

export class Planner extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUserId(): string;
  setUserId(value: string): void;

  getDate(): number;
  setDate(value: number): void;

  getMealtime(): Planner.Mealtime;
  setMealtime(value: Planner.Mealtime): void;

  clearRecipeIdsList(): void;
  getRecipeIdsList(): Array<string>;
  setRecipeIdsList(value: Array<string>): void;
  addRecipeIds(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Planner.AsObject;
  static toObject(includeInstance: boolean, msg: Planner): Planner.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Planner, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Planner;
  static deserializeBinaryFromReader(message: Planner, reader: jspb.BinaryReader): Planner;
}

export namespace Planner {
  export type AsObject = {
    id: string,
    userId: string,
    date: number,
    mealtime: Planner.Mealtime,
    recipeIdsList: Array<string>,
  }

  export enum Mealtime {
    BREAKFAST = 0,
    LUNCH = 1,
    TEA = 2,
    SUPPER = 3,
    SNACK = 4,
  }
}

