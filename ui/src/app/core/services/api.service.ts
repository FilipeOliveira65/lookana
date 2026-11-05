import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { NewServiceInterface } from '../interfaces/new-service';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})

export class ApiService {
  apiUrl = environment.apiUrl;

  constructor(
    private http: HttpClient
  ){}

  newServiceDataToTheApi(serviceData: NewServiceInterface) {
    console.log("Api Service")
    return this.http.post(
      `${this.apiUrl}/create-service`, 
      serviceData
    )
  }
}
