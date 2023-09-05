package com.rydata.springboot3.Pojos;

import java.util.Set;

import com.rydata.springboot3.entities.CourseContents;

public class CourseRequest {

    public int id;

    public String coursename;

    public Set<CourseContents> coursecontents;
    
}
