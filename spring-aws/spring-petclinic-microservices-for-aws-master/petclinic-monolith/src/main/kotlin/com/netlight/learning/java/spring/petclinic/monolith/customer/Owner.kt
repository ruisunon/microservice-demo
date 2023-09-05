/*
 * MIT License
 *
 * Copyright (c) 2020 Alexander Ungar
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package com.netlight.learning.java.spring.petclinic.monolith.customer

import com.fasterxml.jackson.annotation.JsonIgnoreProperties
import com.fasterxml.jackson.annotation.JsonManagedReference
import com.fasterxml.jackson.annotation.JsonView
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import javax.persistence.*
import javax.validation.constraints.Digits
import javax.validation.constraints.NotBlank

@Entity
@Table(name = "owners")
data class Owner(
        @Id
        @GeneratedValue
        val id: Long,
        @NotBlank
        var firstName: String,
        @NotBlank
        var lastName: String,
        @NotBlank
        var address: String,
        @NotBlank
        var city: String,
        @NotBlank
        @Digits(fraction = 0, integer = 10)
        var telephone: String
) {
        // Break circular dependency for auto generated methods
        @OneToMany(cascade = [CascadeType.ALL], mappedBy = "owner")
        @JsonIgnoreProperties("owner")
        var pets: MutableSet<Pet> = mutableSetOf()
}

@Repository
interface OwnerRepository: JpaRepository<Owner, Long>