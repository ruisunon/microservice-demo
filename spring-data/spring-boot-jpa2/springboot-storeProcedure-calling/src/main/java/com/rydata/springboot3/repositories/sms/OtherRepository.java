package com.rydata.springboot3.repositories.sms;

import com.rydata.springboot3.entities.sms.Other;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface OtherRepository extends JpaRepository<Other,Integer> {
    
}
