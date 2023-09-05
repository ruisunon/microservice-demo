package com.rydata.springboot3.services;

import java.util.stream.Collectors;

import com.rydata.springboot3.Pojos.CourseRequest;
import com.rydata.springboot3.entities.Course;
import com.rydata.springboot3.entities.CourseContents;
import com.rydata.springboot3.repositories.CourseContentRepository;
import com.rydata.springboot3.repositories.CourseRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CourseService {

    @Autowired
    CourseRepository courseRepository;

    @Autowired
    CourseContentRepository coursecontentRepo;

    public CourseService(){}

    public Course addCourseWithContents(CourseRequest courseRequest){
        Course course =new Course();
        course.setId(courseRequest.id);
        course.setCoursename(courseRequest.coursename);
        course.setCoursecontents(courseRequest.coursecontents
        .stream()
        .map(coursecontent -> {
            CourseContents ccontents = coursecontent;
            if(ccontents.getId() > 0){
                ccontents = coursecontentRepo.findById(ccontents.getId());
            }
            ccontents.addCourse(course);
            return ccontents;
        })
        .collect(Collectors.toSet()));
        
        return courseRepository.save(course);
    }
    
}
