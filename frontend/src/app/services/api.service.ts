import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { tap } from "rxjs/operators/";
import { environment } from "../../environments/environment";
import { StorageService } from "./storage.service";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  constructor(
    private http: HttpClient,
    private storageService: StorageService
  ) {}

  books() {
    return this.http.get(environment.api + "books/");
  }

  _buildHeaders() {
    const token = this.storageService.getToken();

    const headers = new HttpHeaders({
      'Access-Control-Allow-Headers': 'Authorization',
      'Authorization': "Bearer " + token,
    });

    return headers;
  }

  createBook(title: string, year: number) {
    
    const headers = this._buildHeaders();

    return this.http.post(environment.api + "books/", { title, year }, { headers });
  }

  reviews(bookId: string) {
    return this.http.get(environment.api + `reviews/books/${bookId}`);
  }

  createReview(entityKey: string, entityId: string, content: string) {
    const headers = this._buildHeaders();
    console.log(headers)
    return this.http.post(
      environment.api + "reviews/" + entityKey + "/" + entityId,
      { content },
      { headers }
    );
  }

  signIn(email: string, password: string) {
    return this.http
      .post(environment.api + "authorization/signin", { email, password })
      .pipe(
        tap((response: any) => {
          this.storageService.setToken(response.token);
        })
      );
  }
}
