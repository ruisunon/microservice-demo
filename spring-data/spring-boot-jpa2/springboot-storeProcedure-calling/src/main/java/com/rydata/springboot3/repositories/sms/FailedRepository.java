package com.rydata.springboot3.repositories.sms;

import com.rydata.springboot3.entities.sms.Failed;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FailedRepository extends JpaRepository<Failed,Integer> {
    
}
