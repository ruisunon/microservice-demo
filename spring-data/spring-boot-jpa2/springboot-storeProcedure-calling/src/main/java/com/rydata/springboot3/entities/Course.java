package com.rydata.springboot3.entities;



import java.util.HashSet;
import java.util.Set;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.Table;

import com.fasterxml.jackson.annotation.JsonIgnore;

@Entity
@Table(name ="courses")
public class Course {

    @Id
    @Column(name ="course_id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    private String coursename;

    @ManyToMany(cascade = {CascadeType.DETACH, CascadeType.MERGE, CascadeType.PERSIST, CascadeType.REFRESH})
    @JoinTable(name ="course_coursecontents_mapping",
    joinColumns = @JoinColumn(name ="course_id"),
    inverseJoinColumns = @JoinColumn(name ="content_id"))
    @JsonIgnore
    private Set<CourseContents> coursecontents = new HashSet<>();

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getCoursename() {
        return coursename;
    }

    public void setCoursename(String coursename) {
        this.coursename = coursename;
    }

    public Set<CourseContents> getCoursecontents() {
        return coursecontents;
    }

    public void setCoursecontents(Set<CourseContents> coursecontents) {
        this.coursecontents = coursecontents;
    }

    
    
}
