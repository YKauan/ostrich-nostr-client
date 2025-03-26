import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { FeedData } from './models/Feed.Module';

@Injectable({
  providedIn: 'root'
})
export class GoNostrBackendService {

  apiUrl = 'http://localhost:8080';

  constructor(private httpClient: HttpClient) {}
  
  getFeedData(): Observable<FeedData[]>{
    return this.httpClient.get<FeedData[]>(`${this.apiUrl}/con?npub=npub1xvznjyl57wflha2vm5lql6xu6qc85yx5dymgzrcsjtc3687j3wass4jrvt`);
  }
}
