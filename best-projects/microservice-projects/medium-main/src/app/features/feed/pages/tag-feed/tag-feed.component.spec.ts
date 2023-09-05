import { ComponentFixture, TestBed } from '@angular/core/testing';
import TagFeedComponent from '@feed/pages/tag-feed/tag-feed.component';

describe('TagFeedComponent', () => {
  let component: TagFeedComponent;
  let fixture: ComponentFixture<TagFeedComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [TagFeedComponent],
    });
    fixture = TestBed.createComponent(TagFeedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
