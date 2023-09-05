import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FeedContainerComponent } from './feed-container.component';

describe('FeedContainerComponent', () => {
  let component: FeedContainerComponent;
  let fixture: ComponentFixture<FeedContainerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [FeedContainerComponent]
    });
    fixture = TestBed.createComponent(FeedContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
