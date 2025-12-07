#!/bin/sh

antlr4 -Dlanguage=Go -visitor -package grammar *.g4
