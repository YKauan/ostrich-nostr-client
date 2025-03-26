import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { GoNostrBackendService } from '../../../services/go-nostr-backend.service';
import { PostComponent } from "../post/post.component";
import { FeedData } from '../../../services/models/Feed.Module';

@Component({
  selector: 'app-feed',
  standalone: true,
  imports: [CommonModule, PostComponent],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css'
})
export class FeedComponent implements OnInit {

  public feedData: FeedData[] = [];

  constructor(
    private goNostrBackendService: GoNostrBackendService
  ) {}

  ngOnInit() {
    this.goNostrBackendService.getFeedData().subscribe((feedData) => {
      this.feedData = feedData;
    });
  }
}
