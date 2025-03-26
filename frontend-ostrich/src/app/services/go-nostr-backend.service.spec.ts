import { TestBed } from '@angular/core/testing';

import { GoNostrBackendService } from './go-nostr-backend.service';

describe('GoNostrBackendService', () => {
  let service: GoNostrBackendService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GoNostrBackendService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
