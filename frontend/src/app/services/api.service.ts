import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }

  books(){
    return this.http.get(environment.api + 'books/')
  }

  createBook(title: string, year: number) {
    return this.http.post(environment.api + 'books/', { title, year })
  }

}
