import { Inject, Injectable } from "@angular/core";


@Injectable({
  providedIn: "root",
})
export class StorageService {

  constructor() {}

  setToken(token: string) {
    localStorage.setItem("TOKEN", token);
  }

  getToken(): string {
    return localStorage.getItem("TOKEN");
  }
}
